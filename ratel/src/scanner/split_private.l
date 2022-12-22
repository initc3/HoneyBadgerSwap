%option noyywrap

%{


#include <string>
#include <vector>
#include <assert.h>
#include <set>

using namespace std;

bool finalizeOnChain;

bool isPureMpc, isEvent, isReadDB, isWriteDB, isMpcInput, isMpcOutput, isZkrp; /* see keyword in the current line */
bool inPureMpc; /* in pureMpc block */
bool inFunc; /* in function block */

bool hasFloatPoint;

int braces;

string fileName;
string funcName;
string mpcFile;
FILE *mpcFp, *pythonFp;

int cntMPC;
bool isMPC; /* MP-SPDZ code */
int numInput;

int cntZkrp;
int sumZkrp = 0;

string textLine;

string textFuncUpper; /* read from log, recover private input, acquire locks */
string textFuncLower;

string textPythonUpper; /* only used for monitor function */
string textPythonLower;

string textMpc;

string textRecover;

string returnLine;

bool isArg;
int cntArg;
string arg;
vector<string> args;

string indent;

int cntReadDB;

vector<string> events;
vector<string> eventArgs;

set<string> readKeySet, writeKeySet;


string removeType(string st) {
    for (int i = st.length() - 1; i >= 0; i--) {
        if (st[i] == ' ') {
            return st.substr(i + 1);
        }
    }
}


string getType(string st) {
    for (int i = st.length() - 1; i >= 0; i--) {
        if (st[i] == ' ') {
            return st.substr(0, i);
        }
    }
}


string tab(int repetition) {
    string ret = "";
    for (int i = 0; i < repetition; i++) {
        ret += "    ";
    }
    return ret;
}


string removeIndent(string text) {
    for (int i = 0; i < text.length(); i++) {
        if (text[i] != ' ') {
            return text.substr(i);
        }
    }
}


string getIndent(string text) {
    for (int i = 0; i < text.length(); i++) {
        if (text[i] != ' ') {
            return text.substr(0, i);
        }
    }
}

void checkFloatConstant(){
    int pos_point = 0;
    for(int i = textLine.length()-1;i >= 0;--i){
        if(textLine[i] == '.'){
            pos_point = i;
            break;
        }
    }
    if (pos_point == 0){
        return;
    }

    if(textLine[pos_point - 1] < '0' || textLine[pos_point - 1] > '9'){
        return;
    }
    if(textLine[pos_point + 1] < '0' || textLine[pos_point + 1] > '9'){
        return;
    }

    //get Float Constant
    int pos_l, pos_r;
    for(pos_l = pos_point - 1;pos_l >= 0;--pos_l){
        if(textLine[pos_l] < '0' || textLine[pos_l] > '9'){
            break;
        }
    }
    pos_l = pos_l + 1;

    for(pos_r = pos_point + 1;pos_r < textLine.length();++pos_r){
        if(textLine[pos_r] < '0' || textLine[pos_r] > '9'){
            break;
        }
    }
    pos_r = pos_r - 1;

    
    int lenFloat = pos_r - pos_point;
    int lenInt = pos_point - pos_l;
    int lenNum = pos_r - pos_l + 1;

    string strFloatConstantInt = textLine.substr(pos_l,lenInt) + textLine.substr(pos_point+1,lenFloat);
    int pos_act;
    for(pos_act = 0;pos_act < strFloatConstantInt.length();++pos_act){
        if(strFloatConstantInt[pos_act] != '0'){
            break;
        }
    }
    strFloatConstantInt = strFloatConstantInt.substr(pos_act);

    string strReplace = strFloatConstantInt + " * pow(inv_10," + to_string(lenFloat) + ",prime) % prime";
    
    textLine.replace(pos_l,lenNum,strReplace);
}

void initPython() {
    textPythonUpper += "import asyncio\n";
    textPythonUpper += "import json\n";
    textPythonUpper += "from ratel.src.python.utils import location_sharefile, prog, mpcPort, prime, inv_10, sz, int_to_hex, hex_to_int, recover_input, fp, mark_finish, read_db, write_db, bytes_to_int, bytes_to_list, bytes_to_dict, int_to_bytes, list_to_bytes, dict_to_bytes, execute_cmd, MultiAcquire, sign_and_send, verify_proof\n";
    textPythonUpper += "from ratel.benchmark.src.test_mpc import run_online\n";
    textPythonUpper += "\n";

    textPythonUpper += "async def monitor(server):\n";
    textPythonUpper += "    blkNum = server.web3.eth.get_block_number()\n";
    textPythonUpper += "    while True:\n";
    textPythonUpper += "        curBlkNum = server.web3.eth.get_block_number()\n";
    textPythonUpper += "        if curBlkNum - blkNum > server.confirmation:\n";

    string pythonFile = "python/" + fileName + ".py";
    pythonFp = fopen(pythonFile.c_str(), "w");
}


void finalizePython() {
    textPythonUpper += "            logs = []\n";
    textPythonUpper += "\n";

    events.push_back("GenInputMask");

    if (finalizeOnChain) {
        events.push_back("Finalization");
    }

    for (vector<string>::iterator it = events.begin(); it != events.end(); it++) {
        string eventName = *it;
        textPythonUpper += "            eventFilter = server.contract.events." + eventName + ".createFilter(fromBlock=blkNum, toBlock=curBlkNum - server.confirmation)\n";
        textPythonUpper += "            _logs = eventFilter.get_all_entries()\n";
        textPythonUpper += "            for log in _logs:\n";
        textPythonUpper += "                logs.append((log[\'blockNumber\'], log[\'transactionIndex\'], \'" + eventName + "\', log))\n";
    }
    textPythonUpper += "\n";

    textPythonUpper += "            logs.sort(key=lambda s:(s[0], s[1]))\n";

    textPythonUpper += "            for log in logs:\n";
    textPythonUpper += "                server.loop.create_task(eval(f'run{log[2]}')(server, log[3]))\n";
    textPythonUpper += "            blkNum = curBlkNum - server.confirmation + 1\n";
    textPythonUpper += "        else:\n";
    textPythonUpper += "            await asyncio.sleep(1)\n";
    textPythonUpper += "\n\n";

    if (finalizeOnChain) {
        textPythonUpper += "async def runFinalization(server, log):\n";
        textPythonUpper += "    seq = log['args']['seq']\n";
        textPythonUpper += "    from ratel.genfiles.python." + fileName + "Recover import recover\n";
        textPythonUpper += "    keys = recover(server.contract, seq, 'allSet')\n";
        textPythonUpper += "    for key in keys:\n";
        textPythonUpper += "        server.dbLock[key.lower()].release()\n";
        textPythonUpper += "    mark_finish(server, seq)\n";
        textPythonUpper += "\n\n";
    }

    textPythonUpper += "async def runGenInputMask(server, log):\n";
    textPythonUpper += "    input_mask_cnt = log['args']['inputMaskCnt']\n";
    textPythonUpper += "    input_mask_version = log['args']['inputMaskVersion']\n";
    textPythonUpper += "    await server.gen_input_mask(input_mask_cnt, input_mask_version)\n";
    textPythonUpper += "\n";


    fprintf(pythonFp, "%s", textPythonUpper.c_str());
    fprintf(pythonFp, "%s", textPythonLower.c_str());
    fclose(pythonFp);
}


void initMPC() {
    mpcFile = "mpc/" + fileName + funcName + to_string(++cntMPC) + ".mpc";
    mpcFp = fopen(mpcFile.c_str(), "w");

    textMpc += "bit_length = 210\n";
    textMpc += "\n";

    textMpc += "def read_sint(idx):\n";
    textMpc += "    start = regint(idx)\n";
    textMpc += "    _, shares = sint.read_from_file(start, 1)\n";
    textMpc += "    share = shares[0]\n";
    textMpc += "    return share\n";
    textMpc += "\n";

    textMpc += "def main():\n";
}


void finalizeMPC() {
    textMpc += "main()";

    fprintf(mpcFp, "%s", textMpc.c_str());
    fclose(mpcFp);

    textMpc = "";
}


void processPureMpc() {
    textFuncUpper += "\n";
    textFuncUpper += "async def run" + funcName + "(";

    int num = args.size();
    if (num > 0) textFuncUpper += args[0];
    for (int i = 1; i < num; i++) {
        textFuncUpper += ", " + args[i];
    }

    textFuncUpper += ", port=4900):\n";

    textFuncUpper += "\n";

}


void processEventParam() {
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textLine += ", ";
        if (args[i][0] == '$') {
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textLine += "idx" + postVarName + ", " + "masked" + postVarName;
            if (args[i][1] == '#') {
                textLine += ", idx" + postVarName + "Blinding, masked" + postVarName + "Blinding, proof" + postVarName + ", commitment" + postVarName;
            }
        } else if(getType(args[i]) == "zkp_stmt"){
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textLine += "zkp_stmt" + postVarName + "str, idx" + postVarName + ", masked" + postVarName + ", idx" + postVarName + "Blinding, masked" + postVarName + "Blinding,  proof" + postVarName + ", commitment" + postVarName;
        } else {
            textLine += removeType(args[i]);
        }
    }
}


void monitorEvent() {
    events.push_back(funcName);

    textFuncUpper += "\n";
    textFuncUpper += "async def run" + funcName + "(server, log):\n";

    textFuncUpper += tab(1) + "seq" + funcName + " = log['args']['seq" + funcName + "']\n";
    int num = args.size();
    for (int i = 0; i < num; i++) {
        string arg = args[i];

        string type = getType(arg);
        string varName = removeType(arg);

        if (type[0] == '$') {
            varName[0] = toupper(varName[0]);

            string idxVarName = "idx" + varName;
            textFuncUpper += tab(1) + idxVarName + " = log['args']['" + idxVarName + "']\n";

            string maskedVarName = "masked" + varName;
            textFuncUpper += tab(1) + maskedVarName + " = log['args']['" + maskedVarName + "']\n";

            if (type[1] == '#') {
                string idxVarNameBlinding = "idx" + varName + "Blinding";
                textFuncUpper += tab(1) + idxVarNameBlinding + " = log['args']['" + idxVarNameBlinding + "']\n";

                string maskedVarNameBlinding = "masked" + varName + "Blinding";
                textFuncUpper += tab(1) + maskedVarNameBlinding + " = log['args']['" + maskedVarNameBlinding + "']\n";

                string proofVarName = "proof" + varName;
                textFuncUpper += tab(1) + proofVarName + " = log['args']['" + proofVarName + "']\n";

                string commitmentVarName = "commitment" + varName;
                textFuncUpper += tab(1) + commitmentVarName + " = log['args']['" + commitmentVarName + "']\n";
            }

        } else if (type == "zkp_stmt"){
            varName[0] = toupper(varName[0]);

            string pfstrName = "zkp_stmt" + varName + "str";
            textFuncUpper += tab(1) + pfstrName + " = log['args']['" + pfstrName + "']\n";

            string idxVarName = "idx" + varName;
            textFuncUpper += tab(1) + idxVarName + " = log['args']['" + idxVarName + "']\n";

            string maskedVarName = "masked" + varName;
            textFuncUpper += tab(1) + maskedVarName + " = log['args']['" + maskedVarName + "']\n";

            string idxVarNameBlinding = "idx" + varName + "Blinding";
            textFuncUpper += tab(1) + idxVarNameBlinding + " = log['args']['" + idxVarNameBlinding + "']\n";

            string maskedVarNameBlinding = "masked" + varName + "Blinding";
            textFuncUpper += tab(1) + maskedVarNameBlinding + " = log['args']['" + maskedVarNameBlinding + "']\n";

            string proofVarName = "proof" + varName;
            textFuncUpper += tab(1) + proofVarName + " = log['args']['" + proofVarName + "']\n";

            string commitmentVarName = "commitment" + varName;
            textFuncUpper += tab(1) + commitmentVarName + " = log['args']['" + commitmentVarName + "']\n";
        } else {
            textFuncUpper += tab(1) + varName + " = log['args']['" + varName + "']\n";
        }
    }
    textFuncUpper += "\n";

    eventArgs = args;
}


void unmaskInput() {
    bool flag = 0;
    bool hasZKP = 0;

//    textFuncUpper += "    liszkp = []\n";

    for (int i = 0; i < args.size(); i++) {
        if (args[i][0] == '$') {
            flag = 1;
            string varName = removeType(args[i]);
            string postVarName = varName;
            postVarName[0] = toupper(postVarName[0]);
            textFuncUpper += "    " + varName + " = recover_input(server.db, masked" + postVarName + ", idx" + postVarName + ")\n";
            if (args[i][1] == '#') {
                textFuncUpper += "    " + varName + "Blinding = recover_input(server.db, masked" + postVarName + "Blinding, idx" + postVarName + "Blinding)\n";

                hasZKP = 1;
            }
        } else if(getType(args[i]) == "zkp_stmt"){
            flag = 1;
            string varName = removeType(args[i]);
            string postVarName = varName;
            postVarName[0] = toupper(postVarName[0]);
            textFuncUpper += "    " + varName + " = recover_input(server.db, masked" + postVarName + ", idx" + postVarName + ")\n";
            textFuncUpper += "    " + varName + "Blinding = recover_input(server.db, masked" + postVarName + "Blinding, idx" + postVarName + "Blinding)\n";

            textFuncUpper += "    liszkp.append((zkp_stmt" + postVarName + "str, idx" + postVarName + ", masked" + postVarName + ", idx" + postVarName + "Blinding, masked" + postVarName + "Blinding, proof" + postVarName + ", commitment" + postVarName + ",32))\n";

            hasZKP = 1;
        }
    }
    if (flag) textFuncUpper += "\n";
    if (hasZKP) textFuncLower += "\n";
}


void collectKeys() {
    string tmp;

    textFuncUpper += "    readKeys = [";
    for (set<string>::iterator it = readKeySet.begin(); it != readKeySet.end(); it++) {
        tmp += *it + ", ";
    }

    if (tmp.length() > 0) tmp = tmp.substr(0, tmp.size() - 2);
    textFuncUpper += tmp + "]\n";

    tmp = "";
    textFuncUpper += "    writeKeys = [";
    for (set<string>::iterator it = writeKeySet.begin(); it != writeKeySet.end(); it++) {
        tmp += *it + ", ";
    }
    if (tmp.length() > 0) tmp = tmp.substr(0, tmp.size() - 2);
    textFuncUpper += tmp + "]\n";

    textFuncUpper += "    readKeys =  [k.lower() for k in readKeys]\n";
    textFuncUpper += "    writeKeys =  [k.lower() for k in writeKeys]\n";

    textFuncUpper += "\n";
}


void createLocks() {
    bool flag;

    if (readKeySet.size() > 0) {
        flag = true;
        textFuncUpper += "    for key in readKeys:\n";
        textFuncUpper += "        if key not in server.dbLock.keys():\n";
        textFuncUpper += "            server.dbLock[key] = asyncio.Lock()\n";
        textFuncUpper += "            server.dbLockCnt[key] = 0\n";
    }

    if (writeKeySet.size() > 0) {
        flag = true;
        textFuncUpper += "    for key in writeKeys:\n";
        textFuncUpper += "        if key not in server.dbLock.keys():\n";
        textFuncUpper += "            server.dbLock[key] = asyncio.Lock()\n";
        textFuncUpper += "            server.dbLockCnt[key] = 0\n";
    }

    if (flag) textFuncUpper += "\n";
}


void acquireLocks() {
    textFuncUpper += "    tasks = []\n";

    if (!inPureMpc) textFuncUpper += "    port = mpcPort(seq" + funcName + ", server.concurrency)\n";
    if (finalizeOnChain) {
        textFuncUpper += "    tasks.append(server.portLock[port])\n";

        if (readKeySet.size() > 0) {
            textFuncUpper += "    for key in readKeys:\n";
            textFuncUpper += "        tasks.append(server.dbLock[key])\n";
        }

        if (writeKeySet.size() > 0) {
            textFuncUpper += "    for key in writeKeys:\n";
            textFuncUpper += "        if key not in readKeys:\n";
            textFuncUpper += "            tasks.append(server.dbLock[key])\n";
        }

        textFuncUpper += "    await MultiAcquire(tasks)\n";

    } else {
        textFuncUpper += "    tasks.append(server.portLock[port].acquire())\n";

        if (readKeySet.size() > 0) {
            textFuncUpper += "    for key in readKeys:\n";
            textFuncUpper += "        tasks.append(server.dbLock[key].acquire())\n";
        }

        if (writeKeySet.size() > 0) {
            textFuncUpper += "    for key in writeKeys:\n";
            textFuncUpper += "        if key not in readKeys:\n";
            textFuncUpper += "            tasks.append(server.dbLock[key].acquire())\n";
        }

        textFuncUpper += "    await asyncio.wait(tasks)\n";
    }
    textFuncUpper += "\n";

    bool flag;

    if (readKeySet.size() > 0) {
        flag = true;
        textFuncUpper += "    for key in readKeys:\n";
        textFuncUpper += "        server.dbLockCnt[key] += 1\n";
    }

    if (writeKeySet.size() > 0) {
        flag = true;
        textFuncUpper += "    for key in writeKeys:\n";
        textFuncUpper += "        server.dbLockCnt[key] += 1\n";
    }

    if (flag) textFuncUpper += "\n";
}


void setRecover() {
    if (eventArgs.size()) {
        string parameter = "(seq" + funcName;
        int num = eventArgs.size();
        for (int i = 0; i < num; i++) {
            string arg = eventArgs[i];

            string type = getType(arg);
            string varName = removeType(arg);

            if (type[0] == '$') {
                varName[0] = toupper(varName[0]);

                string idxVarName = "idx" + varName;
                parameter += ", " + idxVarName;

                string maskedVarName = "masked" + varName;
                parameter += ", " + maskedVarName;

            } else {
                parameter += ", " + varName;
            }
        }
        parameter += ")";

        textRecover += "def writeSet" + funcName + parameter + ":\n";
        textRecover += "    return set({";
        for (set<string>::iterator it = writeKeySet.begin(); it != writeKeySet.end(); it++) {
            textRecover += *it + ", ";
        }
        textRecover += "})\n";
        textRecover += "\n\n";

        textRecover += "def readSet" + funcName + parameter + ":\n";
        textRecover += "    return set({";
        for (set<string>::iterator it = readKeySet.begin(); it != readKeySet.end(); it++) {
            textRecover += *it + ", ";
        }
        textRecover += "})\n";
        textRecover += "\n\n";

        textRecover += "def allSet" + funcName + parameter + ":\n";
        textRecover += "    return writeSet" + funcName + parameter + ".union(readSet" + funcName + parameter + ")\n";
        textRecover += "\n\n";
    }
}


void readDB() {
    string textTmp = textLine;

    string value = "value" + to_string(++cntReadDB);
    string key = args[0];
    string type = args[1];

    readKeySet.insert(key);

    textLine = indent + value + " = read_db(server, " + key + ", " + char('0' + int(finalizeOnChain)) + ")\n";
    if (type == "int") {
        textLine += indent + value + " = bytes_to_int(" + value + ")\n";
    } else if (type == "list") {
        textLine += indent + value + " = bytes_to_list(" + value + ")\n";
    } else if (type == "dict") {
        textLine += indent + value + " = bytes_to_dict(" + value + ")\n";
    } else {
        assert(0);
    }
    textLine += textTmp + value;
}


void writeDB() {
    string key = args[0];
    string value = args[1];
    string type = args[2];

    writeKeySet.insert(key);

    if (type == "int") {
        textLine += indent + value + " = int_to_bytes(" + value + ")\n";
    } else if (type == "list") {
        textLine += indent + value + " = list_to_bytes(" + value + ")\n";
    } else if (type == "dict") {
        textLine += indent + value + " = dict_to_bytes(" + value + ")\n";
    } else {
        assert(0);
    }

    textLine += indent + "write_db(server, " + key + ", " + value + ", " + char('0' + int(finalizeOnChain)) + ")";

}


void mpcInput() {
    /* python */
    textFuncLower += indent + "file = location_sharefile(server.serverID, port)\n";
    textFuncLower += indent + "with open(file, \"wb\") as f:\n";
    textFuncLower += indent + "    f.write(\n";
    numInput = args.size();
    for (int i = 0; i < numInput; i++) {
        textFuncLower += indent + tab(2) + ((i == 0) ? "" : "+ ") + "int_to_hex(" + removeType(args[i]) + ")\n";
    }
    textFuncLower += indent + tab(1) + ")\n";
    textFuncLower += "\n";

    /* MP-SPDZ */
    initMPC();
    for (int i = 0; i < numInput; i++) {
        string varType = getType(args[i]);
        string varName = removeType(args[i]);
        if (varType == "sint") {
            textMpc += tab(1) + varName + " = read_sint(" + to_string(i) + ")\n";
        } else {
            assert(varType == "sfix");
            textMpc += tab(1) + varName + " = sfix._new(read_sint(" + to_string(i) + "))\n";
        }
    }

    /* python */
    string mpcName = mpcFile.substr(4, mpcFile.length() - 8);
    textFuncLower += indent + "await run_online(server.serverID, port, server.players, server.threshold, \'" + mpcName + "\', seq" + funcName + ")\n";
    textFuncLower += "\n";
}


void mpcOutput() {
    /* MP-SPDZ */
    int num = args.size();
    for (int i = 0; i < num; i++) {
        string varType = getType(args[i]);
        string varName = removeType(args[i]);
        if (varType == "sfix") {
            textMpc += tab(1) + varName + " = " + varName + ".v\n";
        } else if (varType == "cint") {
            textMpc += tab(1) + varName + " = sint(" + varName + ")\n";
        } else if (varType == "cfix") {
            textMpc += tab(1) + varName + " = sint(" + varName + ".v)\n";
        } else {
            assert(varType == "sint");
        }
        textMpc += tab(1) + "sint.write_to_file(" + varName + ")\n";
    }
    textMpc += "\n";

    finalizeMPC();

    /* python */
    textFuncLower += indent + "input_arg_num = " + to_string(numInput) + "\n";
    textFuncLower += indent + "with open(file, \"rb\") as f:\n";
    textFuncLower += indent + "    f.seek(input_arg_num * sz)\n";
    for (int i = 0; i < num; i++) {
        textFuncLower += indent + tab(1) + removeType(args[i]) + " = hex_to_int(f.read(sz))\n";
    }
}

void zkrp(){
    if(cntZkrp == 1){
        textFuncLower += "    lis_zkpstmt = json.loads(json_lis_zkpstmt)\n";
    }

    string pfstr = args[0];

    textFuncLower += "    print(\'" + pfstr +  "\')\n";

    

    int pospf = 0;
    for(int i = 0;i < pfstr.length();++i){
        if(pfstr[i] == '<' || pfstr[i] == '>'){
            pospf = i;
            break;
        }
    }

    string pfVarName = "pfVar" + to_string(cntZkrp-1);

    string pfExp = pfstr.substr(0,pospf);
    string zkpstmtName = "zkpstmt" + to_string(cntZkrp-1);

    bool hasMul = false;
    int posMul = 0;
    string expL, expR;
    for(int i = 0;i < pfExp.length();++i){
        if(pfExp[i] == '*'){
            hasMul = true;
            posMul = i;
        }
    }

    textFuncLower += "    tmpMul = False\n";
    textFuncLower += "    tmpMulR = 0\n";
    textFuncLower += "    tmptype_Mul = 0\n";

    if(pfstr[pospf] == '>' && pfstr[pospf+1] == '=') {
        string rangeName = pfstr.substr(pospf+2);
        textFuncLower += "    tmpRange = int(" + rangeName + ")\n";

        // secret_value >= r <==> secret_value - r >= 0
        textFuncLower += "    " + pfVarName + " = " + pfExp + " - tmpRange\n";

        if(hasMul){
            textFuncLower += "    " + pfVarName + " = " + pfExp.substr(0,posMul) + ")\n";
            textFuncLower += "    tmpMulR = (" + pfExp.substr(posMul+1) + "\n";
            textFuncLower += "    tmptype_Mul = 1\n";
        }
    } else if(pfstr[pospf] == '>' && pfstr[pospf+1] != '=') {
        string rangeName = pfstr.substr(pospf+1);
        textFuncLower += "    tmpRange = int(" + rangeName + ")\n";

        // secret_value > r <==> secret_value - r - 1 >= 0
        textFuncLower += "    " + pfVarName + " = " + pfExp + " - tmpRange - 1\n";

        if(hasMul){
            textFuncLower += "    " + pfVarName + " = " + pfExp.substr(0,posMul) + ")\n";
            textFuncLower += "    tmpMulR = (" + pfExp.substr(posMul+1) + "\n";
            textFuncLower += "    tmptype_Mul = 2\n";
            textFuncLower += "    tmpRange = tmpRange + 1\n";
        }
    } else if(pfstr[pospf] == '<' && pfstr[pospf+1] == '=') { 
        string rangeName = pfstr.substr(pospf+2);
        textFuncLower += "    tmpRange = int(" + rangeName + ")\n";

        // secret_value <= r <==> r - secret_value >= 0 
        textFuncLower += "    " + pfVarName + " = tmpRange - " + pfExp + "\n";

        if(hasMul){
            textFuncLower += "    " + pfVarName + " = " + pfExp.substr(0,posMul) + ")\n";
            textFuncLower += "    tmpMulR = (" + pfExp.substr(posMul+1) + "\n";
            textFuncLower += "    tmptype_Mul = 3\n";
            textFuncLower += "    tmpRange = -tmpRange\n";
        }
    } else if(pfstr[pospf] == '<' && pfstr[pospf+1] != '=') { 
        string rangeName = pfstr.substr(pospf+1);
        textFuncLower += "    tmpRange = int(" + rangeName + ")\n";

        // secret_value < r <==> r - secret_value - 1 >= 0
        textFuncLower += "    " + pfVarName + " = tmpRange - " + pfExp + " - 1\n";

        if(hasMul){
            textFuncLower += "    " + pfVarName + " = " + pfExp.substr(0,posMul) + ")\n";
            textFuncLower += "    tmpMulR = (" + pfExp.substr(posMul+1) + "\n";
            textFuncLower += "    tmptype_Mul = 4\n";
            textFuncLower += "    tmpRange = -tmpRange + 1\n";
        }
    } else {
        assert(0);
    }

    textFuncLower += "    " + zkpstmtName + " = lis_zkpstmt[" + to_string(cntZkrp-1) + "]\n";
//    textFuncLower += "    print(\'cntZkrp:" + to_string(cntZkrp-1) + "\')\n";
//    textFuncLower += "    print(\'zkpstmt\'," + zkpstmtName + ")\n";
    textFuncLower += "    print(\'tmpRange:\',tmpRange)\n";
}


void addText(char *yytext) {
    string text = string(yytext);

    if (isArg) {
        arg += text;
    } else if (returnLine.length() > 0) {
        returnLine += text;
    }else {
        textLine += text;
    }
}

void addZkrpArgs(){
    if(cntZkrp > 0){
        string recover_str = "";
        string liszkp_str = "";
        for(int i = 0;i < cntZkrp;++i){
            string postVarName = "zkpstmt" + to_string(i);

            //todos
            /////////////////////add recover input for zkp/////////////////////
//            recover_str += "    " + postVarName + " = recover_input(server.db, masked" + postVarName + ", idx" + postVarName + ")\n";
//            recover_str += "    " + postVarName + "Blinding = recover_input(server.db, masked" + postVarName + "Blinding, idx" + postVarName + "Blinding)\n";
        }
    }
}


void zkrpStmt(){
    string pfVarName = "pfVar" + to_string(cntZkrp);

    string zkpstmtName = "zkpstmt" + to_string(cntZkrp);

    textLine += "await verify_proof(server, " + pfVarName + ", " + zkpstmtName + " , tmptype_Mul, tmpMulR, tmpRange)";
}



%}


SPACE [ \t\r]+


%%


.*import.*\n {
    string text = string(yytext);
    for (int i = 0; i < text.length(); i++) {
        if (text[i] == 'i') {
            fprintf(pythonFp, "%s", text.substr(i).c_str());
            break;
        }
    }
}


^pureMpc{SPACE} {
    inPureMpc = true;
    isPureMpc = true;
    isArg = true;
    cntArg = 1;
    cntMPC = 0;
    cntZkrp = 0;
    cntReadDB = 0;
    textFuncUpper = "";
    textFuncLower = "";
    readKeySet.clear();
    writeKeySet.clear();
}


^mpc{SPACE} {
    inFunc = true;
    isEvent = true;
    isArg = true;
    cntArg = 1;
    cntMPC = 0;
    cntZkrp = 0;
    cntReadDB = 0;
    textFuncUpper = "";
    textFuncLower = "";
    readKeySet.clear();
    writeKeySet.clear();
}


^.*readDB\( {
    isReadDB = true;
    isArg = true;
    cntArg = 1;

    string text = string(yytext);
    indent = getIndent(text);
    textLine += text.substr(0, text.length() - strlen("readDB("));
}


^.*writeDB\( {
    isWriteDB = true;
    isArg = true;
    cntArg = 1;

    string text = string(yytext);
    indent = getIndent(text);
}


^.*mpcInput\( {
    isMpcInput = true;
    isArg = true;
    cntArg = 1;

    indent = getIndent(string(yytext));
}


^.*mpcOutput\( {
    isMpcOutput = true;
    isArg = true;
    cntArg = 1;
}

^.*zkrp\( {
    isZkrp = true;
    isArg = true;
    cntArg = 1;

    indent = getIndent(string(yytext));
    string text = string(yytext);
    textLine += text.substr(0, text.length() - strlen("zkrp("));
}

^.*return {
    returnLine = string(yytext);
}


\( {
    if(isArg){
        cntArg++;
    }
    if (isEvent || isPureMpc) {
        funcName = arg;
        arg = "";
    } else {
        addText(yytext);
    }
}


\) {
    if (isArg) {
        if (arg.length()) {
            if(!isZkrp){
                args.push_back(arg);
                arg = "";
            } else if(cntArg == 1){
                args.push_back(arg);
                arg = "";
            } else {
                addText(yytext);
            }
        }

        if (isReadDB) {
            assert(args.size() == 2);
            readDB();
        } else if (isWriteDB) {
            assert(args.size() == 3);
            writeDB();
        } else if (isZkrp){
            if(cntArg == 1){
                zkrpStmt();
            }
        }
        
        cntArg--;
        if(!isZkrp){
            isArg = false;
        } else if(cntArg == 0){
            isArg = false;
        }
    } else {
        addText(yytext);
    }
}


,{SPACE}* {
    if (isArg) {
        args.push_back(arg);
        arg = "";
    } else {
        addText(yytext);
    }
}


\{ {
    braces++;
    addText(yytext);
}


\} {
    braces--;
    if (braces > 0) addText(yytext);
    else {

        collectKeys();
        createLocks();
        acquireLocks();

        setRecover();

        addZkrpArgs();

        textPythonLower += textFuncUpper + textFuncLower;

        if (!inPureMpc) {
            if (finalizeOnChain) {
                textPythonLower += "    tx = server.contract.functions.finalize(seq" + funcName + ").buildTransaction({'from': server.account.address, 'gas': 1000000, 'nonce': server.web3.eth.get_transaction_count(server.account.address)})\n";
                textPythonLower += "    sign_and_send(tx, server.web3, server.account)";
            } else {
                textPythonLower += "    mark_finish(server, seq" + funcName + ")";
            }
        } else {
            textPythonLower += "    server.portLock[port].release()\n";
        }

        textPythonLower += returnLine;

        sumZkrp += cntZkrp;

        inPureMpc = false;
        inFunc = false;
        returnLine = "";
    }
}


\n {
    textLine += "\n";

//    printf("-textLine:%s", textLine.c_str());

    if(hasFloatPoint && (!isMPC)){
        checkFloatConstant();
        hasFloatPoint = false;
    }


    if (isPureMpc) {
        processPureMpc();

        isPureMpc = false;
    } else if (isEvent) {
        monitorEvent();
        unmaskInput();

        isEvent = false;
    } else if (isReadDB) {
        textFuncLower += textLine;

        isReadDB = false;
    } else if (isMpcInput) {
        isMPC = true;
        mpcInput();

        isMpcInput = false;
    } else if (isMpcOutput) {
        mpcOutput();
        isMPC = false;

        isMpcOutput = false;
    } else if(isZkrp){
        cntZkrp++;
        zkrp();
        textFuncLower += textLine;

        isZkrp = false;
    }else if (isWriteDB) {
        textFuncLower += textLine;

        isWriteDB = false;
    } else if (returnLine.length() == 0) {
        if (isMPC) {
            textMpc += tab(1) + removeIndent(textLine);
        } else if (inFunc || inPureMpc) {
            textFuncLower += textLine;
        } else {
            textPythonLower += textLine;
        }
    }

    args.clear();

    textLine = "";
}

\. {
//    hasFloatPoint = true;
    addText(yytext);
}

. {
    addText(yytext);
}


%%


int main(int argc, char **argv) {
    char *_fileName = argv[1];
    fileName = string(_fileName);

    finalizeOnChain = argv[2][0] - '0';

    initPython();

    yylex();

    finalizePython();

    textRecover += "\n";
    FILE *recoverFp;
    string recoverFile = "python/" + fileName + "Recover.py";
    recoverFp = fopen(recoverFile.c_str(), "a");
    fprintf(recoverFp, "%s", textRecover.c_str());
    fclose(recoverFp);

//    FILE *tmpFp;
//    string tmpFile = "tmp.txt";
//    tmpFp = fopen(tmpFile.c_str(), "a");
//    fprintf(tmpFp, "%s", to_string(sumZkrp).c_str());
//    fclose(tmpFp);

	return 0;
}

