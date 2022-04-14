%option noyywrap

%{


#include <string>
#include <vector>
#include <assert.h>
#include <set>

using namespace std;

bool isPureMpc, isEvent, isReadDB, isWriteDB, isMpcInput, isMpcOutput;

int braces;

string fileName;
string funcName;
string mpcFile;
FILE *mpcFp, *pythonFp;

int cntMPC;
bool isMPC; /* MP-SPDZ code */
int numInput;

string textLine;

string textFuncUpper; /* read from log, recover private input, acquire locks */
string textFuncLower;

string textPythonUpper; /* only used for monitor function */
string textPythonLower;

string textMpc;

string textRecover;

bool isArg;
string arg;
vector<string> args;

string indent;

int cntReadDB;

vector<string> events;

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


void initPython() {
    textPythonUpper += "import asyncio\n";
    textPythonUpper += "from ratel.src.python.utils import location_sharefile, prog, mpcPort, blsPrime, sz, int_to_hex, hex_to_int, recover_input, fp, replay, mark_finish, read_db, write_db, bytes_to_int, bytes_to_list, bytes_to_dict, int_to_bytes, list_to_bytes, dict_to_bytes\n";
    textPythonUpper += "\n";

    textPythonUpper += "async def monitor(server, loop):\n";
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
    textPythonUpper += "                for _ in range((1 if log[2] != 'Trade' else replay)):\n";
    textPythonUpper += "                    loop.create_task(eval(f'run{log[2]}')(server, log[3]))\n";
    textPythonUpper += "            blkNum = curBlkNum - server.confirmation + 1\n";
    textPythonUpper += "        else:\n";
    textPythonUpper += "            await asyncio.sleep(1)\n";
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

    textFuncUpper += "):\n";

    textFuncUpper += "    seq" + funcName + " = -1\n";
    textFuncUpper += "\n";

    textRecover += "def recover" + funcName + "():\n";
    textRecover += "    return [";
}


void processEventParam() {
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textLine += ", ";
        if (args[i][0] == '$') {
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textLine += "idx" + postVarName + ", " + "masked" + postVarName;
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

        } else {
            textFuncUpper += tab(1) + varName + " = log['args']['" + varName + "']\n";
        }
    }
    textFuncUpper += "\n";

    textRecover += "def recover" + funcName + "(seq" + funcName;
    for (int i = 0; i < num; i++) {
        string arg = args[i];

        string type = getType(arg);
        string varName = removeType(arg);

        if (type[0] == '$') {
            varName[0] = toupper(varName[0]);

            string idxVarName = "idx" + varName;
            textRecover += ", " + idxVarName;

            string maskedVarName = "masked" + varName;
            textRecover += ", " + maskedVarName;

        } else {
            textRecover += ", " + varName;
        }
    }
    textRecover += "):\n";
    textRecover += "    return [";
}


void unmaskInput() {
    bool flag = 0;
    for (int i = 0; i < args.size(); i++) {
        if (args[i][0] == '$') {
            flag = 1;
            string varName = removeType(args[i]);
            string postVarName = varName;
            postVarName[0] = toupper(postVarName[0]);
            textFuncUpper += "    " + varName + " = recover_input(server.db, masked" + postVarName + ", idx" + postVarName + ")\n";
        }
    }
    if (flag) textFuncUpper += "\n";
}


void collectKeys() {
    textFuncUpper += "    readKeys = [";
    string tmp;
    for (set<string>::iterator it = readKeySet.begin(); it != readKeySet.end(); it++) {
        tmp += *it + ", ";
    }

    if (tmp.length() > 0) tmp = tmp.substr(0, tmp.size() - 2);
    textFuncUpper += tmp + "]\n";

    textFuncUpper += "    writeKeys = [";
    tmp = "";
    for (set<string>::iterator it = writeKeySet.begin(); it != writeKeySet.end(); it++) {
        tmp += *it + ", ";
    }
    if (tmp.length() > 0) tmp = tmp.substr(0, tmp.size() - 2);
    textFuncUpper += tmp + "]\n";

    textFuncUpper += "\n";
}


void createLocks() {
    textFuncUpper += "    for key in readKeys:\n";
    textFuncUpper += "        if key not in server.dbLock.keys():\n";
    textFuncUpper += "            server.dbLock[key] = asyncio.Lock()\n";
    textFuncUpper += "            server.dbLockCnt[key] = 0\n";

    textFuncUpper += "    for key in writeKeys:\n";
    textFuncUpper += "        if key not in server.dbLock.keys():\n";
    textFuncUpper += "            server.dbLock[key] = asyncio.Lock()\n";
    textFuncUpper += "            server.dbLockCnt[key] = 0\n";

    textFuncUpper += "\n";
}


void acquireLocks() {
    textFuncUpper += "    tasks = []\n";

    textFuncUpper += "    for key in readKeys:\n";
    textFuncUpper += "        tasks.append(server.dbLock[key].acquire())\n";

    textFuncUpper += "    for key in writeKeys:\n";
    textFuncUpper += "        if key not in readKeys:\n";
    textFuncUpper += "            tasks.append(server.dbLock[key].acquire())\n";

    textFuncUpper += "    port = mpcPort(seq" + funcName + ", server.concurrency)\n";
    textFuncUpper += "    tasks.append(server.portLock[port].acquire())\n";

    textFuncUpper += "    await asyncio.wait(tasks)\n";

    textFuncUpper += "\n";

    textFuncUpper += "    for key in readKeys:\n";
    textFuncUpper += "        server.dbLockCnt[key] += 1\n";
    textFuncUpper += "    for key in writeKeys:\n";
    textFuncUpper += "        server.dbLockCnt[key] += 1\n";

    textFuncUpper += "\n";
}


void readDB() {
    string textTmp = textLine;

    string value = "value" + to_string(++cntReadDB);
    string key = args[0];
    string type = args[1];

    readKeySet.insert(key);

    textLine = indent + value + " = read_db(server, " + key + ")\n";
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

    textLine += indent + "write_db(server, " + key + ", " + value + ")";

    textRecover += key + ", ";
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
    textFuncLower += indent + "cmd = f'{prog} -N {server.players} -T {server.threshold} -p {server.serverID} -pn {port} -P {blsPrime} " + mpcFile.substr(4, mpcFile.length() - 8) + "\'\n";
    textFuncLower += indent + "proc = await asyncio.create_subprocess_shell(cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE)\n";
    textFuncLower += indent + "stdout, stderr = await proc.communicate()\n";
    textFuncLower += indent + "print(f'[{cmd!r} exited with {proc.returncode}]')\n";
    textFuncLower += indent + "if stdout:\n";
    textFuncLower += indent + "    print(f'[stdout]\\n{stdout.decode()}')\n";
    textFuncLower += indent + "if stderr:\n";
    textFuncLower += indent + "    print(f'[stderr]\\n{stderr.decode()}')\n";
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


void addText(char *yytext) {
    string text = string(yytext);

    if (isArg) {
        arg += text;
    } else {
        textLine += text;
    }
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
    isPureMpc = true;
    isArg = true;
    cntMPC = 0;
    cntReadDB = 0;
    textFuncUpper = "";
    textFuncLower = "";
    readKeySet.clear();
    writeKeySet.clear();
}


^mpc{SPACE} {
    isEvent = true;
    isArg = true;
    cntMPC = 0;
    cntReadDB = 0;
    textFuncUpper = "";
    textFuncLower = "";
    readKeySet.clear();
    writeKeySet.clear();
}


^.*readDB\( {
    isReadDB = true;
    isArg = true;

    string text = string(yytext);
    indent = getIndent(text);
    textLine += text.substr(0, text.length() - strlen("readDB("));
}


^.*writeDB\( {
    isWriteDB = true;
    isArg = true;

    string text = string(yytext);
    indent = getIndent(text);
}


^.*mpcInput\( {
    isMpcInput = true;
    isArg = true;

    indent = getIndent(string(yytext));
}


^.*mpcOutput\( {
    isMpcOutput = true;
    isArg = true;
}


\( {
    if (isEvent || isPureMpc) {
        funcName = arg;
        arg = "";
    } else {
        addText(yytext);
    }
}


\) {
    if (isArg) {
        args.push_back(arg);
        arg = "";

        if (isReadDB) {
            assert(args.size() == 2);
            readDB();
        } else if (isWriteDB) {
            assert(args.size() == 3);
            writeDB();
        }

        isArg = false;
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
        textRecover += "]\n\n";

        collectKeys();
        createLocks();
        acquireLocks();

        textPythonLower += textFuncUpper + textFuncLower;

        textPythonLower += "\n";
        textPythonLower += "    server.portLock[port].release()\n";
        textPythonLower += "    mark_finish(server, seq" + funcName + ")\n";
        textPythonLower += "\n";
    }
}


\n {
    textLine += "\n";

//    printf("textLine:%s", textLine.c_str());

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
    } else if (isWriteDB) {
        textFuncLower += textLine;

        isWriteDB = false;
    } else {
        if (isMPC) {
            textMpc += tab(1) + removeIndent(textLine);
        } else {
            textFuncLower += textLine;
        }
    }

    args.clear();

    textLine = "";
}


. {
    addText(yytext);
}


%%


int main(int argc, char **argv) {
    char *_fileName = argv[1];
    fileName = string(_fileName);

    initPython();

    yylex();

    finalizePython();

    textRecover += "\n";
    FILE *recoverFp;
    string recoverFile = "python/" + fileName + "Recover.py";
    recoverFp = fopen(recoverFile.c_str(), "a");
    fprintf(recoverFp, "%s", textRecover.c_str());
    fclose(recoverFp);

	return 0;
}
