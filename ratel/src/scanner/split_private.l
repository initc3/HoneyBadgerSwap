%option noyywrap

%{
#include <string>
#include <vector>
#include <assert.h>

using namespace std;

bool isEvent, isReadDB, isWriteDB, isMpcInput, isMpcOutput, isEmptyLine = true;

int braces;
int numInput;
int cntValueDB;

string fileName;
string funcName;
string mpcFile;
FILE *mpcFp, *pythonFp;

int cntMPC;
bool isMPC;

string textLine;
string textPythonUpper, textPythonLower, textMpc;

bool isArg;
string arg;
vector<string> args;

string prefix;

string textRecover;

vector<string> events;

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

string removeSpace(string text) {
    for (int i = 0; i < text.length(); i++) {
        if (text[i] != ' ') {
            return text.substr(i);
        }
    }
}

string getPrefix(string text) {
    for (int i = 0; i < text.length(); i++) {
        if (text[i] != ' ') {
            return text.substr(0, i);
        }
    }
}

void initPython() {
    textPythonUpper += "import ast\n";
    textPythonUpper += "import asyncio\n";
    textPythonUpper += "from ratel.src.python.utils import location_sharefile, prog, mpcPort, blsPrime, sz, int_to_hex, hex_to_int, recover_input, fp\n";
    textPythonUpper += "\n";

    textPythonUpper += "async def monitor(server, loop):\n";
    textPythonUpper += "    blkNum = server.web3.eth.get_block_number()\n";
    textPythonUpper += "    while True:\n";
    textPythonUpper += "        await asyncio.sleep(5)\n";
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
    textPythonUpper += "                asyncio.ensure_future(eval(f'run\{log[2]\}')(server, log[3]), loop=loop)\n";
    textPythonUpper += "\n";

    textPythonUpper += "            blkNum = curBlkNum - server.confirmation + 1\n";
    textPythonUpper += "\n";

    fprintf(pythonFp, "%s", textPythonUpper.c_str());
    fprintf(pythonFp, "%s", textPythonLower.c_str());
    fclose(pythonFp);
}

void initMPC() {
    mpcFile = "mpc/" + fileName + funcName + to_string(++cntMPC) + ".mpc";
    mpcFp = fopen(mpcFile.c_str(), "w");

    textMpc += "bit_length = 212\n";
    textMpc += "\n";

    textMpc += "def read_sint(idx):\n";
    textMpc += "    start = regint(idx * 32)\n";
    textMpc += "    end = regint()\n";
    textMpc += "    share = sint()\n";
    textMpc += "    readsharesfromfile(start, end, share)\n";
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

void monitorEvent() {
    events.push_back(funcName);

    textPythonLower += "\n";
    textPythonLower += "async def run" + funcName + "(server, log):\n";

    textPythonLower += tab(1) + "seq" + funcName + " = log['args']['seq" + funcName + "']\n";
    int num = args.size();
    for (int i = 0; i < num; i++) {
        string arg = args[i];
        textPythonLower += tab(1) + arg + " = log['args']['" + arg + "']\n";
    }
    textPythonLower += "\n";

    textRecover += "def recover" + funcName + "(seq" + funcName;
    for (int i = 0; i < num; i++) {
        textRecover += ", " + args[i];
    }
    textRecover += "):\n";
    textRecover += "    return [";
}

void readDB() {
    string textTmp = textLine;

    string value = "value" + to_string(++cntValueDB);

    textLine = prefix + "try:\n";
    textLine += prefix + "    " + value + " = server.db.Get(" + args[0] + ".lower().encode())\n";
    textLine += prefix + "except KeyError:\n";
    textLine += prefix + "    " + value + " = bytes(0)\n";
    if (args[1] == "int") {
        textLine += prefix + value + " = int.from_bytes(" + value + ", 'big')\n";
    } else if (args[1] == "list") {
        textLine += prefix + "try:\n";
        textLine += prefix + "    " + value + " = " + value + ".decode(encoding='utf-8')\n";
        textLine += prefix + "    " + value + " = list(ast.literal_eval(" + value + "))\n";
        textLine += prefix + "except:\n";
        textLine += prefix + "    " + value + " = []\n";
    } else if (args[1] == "dict") {
        textLine += prefix + "try:\n";
        textLine += prefix + "    " + value + " = " + value + ".decode(encoding='utf-8')\n";
        textLine += prefix + "    " + value + " = dict(ast.literal_eval(" + value + "))\n";
        textLine += prefix + "except:\n";
        textLine += prefix + "    " + value + " = {}\n";
    } else {
        assert(0);
    }
    textLine += textTmp + value;
}

void writeDB() {
    if (args[2] == "int") {
        string key = args[0];
        string value = args[1];
        textLine += "server.db.Put(" + key + ".lower().encode(), " + value + ".to_bytes((" + value + ".bit_length() + 7) // 8, 'big'))";
    } else if (args[2] == "list" || args[2] == "dict") {
        string key = args[0];
        string value = args[1];
        string textTmp;

        textTmp = prefix + value + " = str(" + value + ")\n";
        textTmp += prefix + value + " = bytes(" + value + ", encoding='utf-8')\n";
        textLine += "server.db.Put(" + key + ".lower().encode(), " + value + ")";
        textLine = textTmp + textLine;
    } else {
        assert(0);
    }

    textRecover += args[0] + ", ";
}

void mpcInput() {
    //python
    textPythonLower += "\n";
    textPythonLower += prefix + "file = location_sharefile(server.serverID)\n";
    textPythonLower += prefix + "with open(file, \"wb\") as f:\n";
    textPythonLower += prefix + "    f.write(\n";
    numInput = args.size();
    for (int i = 0; i < numInput; i++) {
        textPythonLower += prefix + tab(2) + ((i == 0) ? "" : "+ ") + "int_to_hex(" + removeType(args[i]) + ")\n";
    }
    textPythonLower += prefix + tab(1) + ")\n";
    textPythonLower += "\n";

    //mpc
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
    textMpc += "\n";

    //python
    textPythonLower += prefix + "port = mpcPort(seq" + funcName + ")\n";
    textPythonLower += prefix + "await server.portLock[port].acquire()\n";
    textPythonLower += prefix + "cmd = f'{prog} -N {server.players} -T {server.threshold} -p {server.serverID} -pn {port} -P {blsPrime} " + mpcFile.substr(4, mpcFile.length() - 8) + "\'\n";
    textPythonLower += prefix + "proc = await asyncio.create_subprocess_shell(cmd, stdout=asyncio.subprocess.PIPE, stderr=asyncio.subprocess.PIPE)\n";
    textPythonLower += prefix + "stdout, stderr = await proc.communicate()\n";
    textPythonLower += prefix + "print(f'[{cmd!r} exited with {proc.returncode}]')\n";
    textPythonLower += prefix + "if stdout:\n";
    textPythonLower += prefix + "    print(f'[stdout]\\n{stdout.decode()}')\n";
    textPythonLower += prefix + "if stderr:\n";
    textPythonLower += prefix + "    print(f'[stderr]\\n{stderr.decode()}')\n";
    textPythonLower += prefix + "server.portLock[port].release()\n";
    textPythonLower += "\n";
}

void mpcOutput() {
    //mpc
    textMpc += "\n";
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
        textMpc += tab(1) + "writesharestofile(" + varName + ")\n";
    }
    textMpc += "\n";

    finalizeMPC();

    //python
    textPythonLower += prefix + "input_arg_num = " + to_string(numInput) + "\n";
    textPythonLower += prefix + "with open(file, \"rb\") as f:\n";
    textPythonLower += prefix + "    f.seek(input_arg_num * sz)\n";
    for (int i = 0; i < num; i++) {
        textPythonLower += prefix + tab(1) + removeType(args[i]) + " = hex_to_int(f.read(sz))\n";
    }
    textPythonLower += "\n";
}

void addText(char *yytext) {
    string text = string(yytext);

    for (int i = 0; i < text.length(); i++) {
        char ch = text[i];
        if (ch != ' ' && ch != '\r' && ch != '\t') {
            isEmptyLine = false;
        }
    }

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

^mpc{SPACE} {
    isEvent = true;
    isArg = true;
    cntMPC = 0;
    cntValueDB = 0;
}

^.*readDB\( {
    isReadDB = true;
    isArg = true;

    string text = string(yytext);
    prefix = getPrefix(text);
    textLine += text.substr(0, text.length() - 7);
}

^.*writeDB\( {
    isWriteDB = true;
    isArg = true;

    string text = string(yytext);
    prefix = getPrefix(text);
    textLine += text.substr(0, text.length() - 8);
}

^.*mpcInput\( {
    isMpcInput = true;
    isArg = true;

    prefix = getPrefix(string(yytext));
}

^.*mpcOutput\( {
    isMpcOutput = true;
    isArg = true;
}

\( {
    if (isEvent) {
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
    else textRecover += "]\n\n";
}

\n {
    textLine += "\n";

//    printf("textLine:%s", textLine.c_str());

    if (isEvent) {
        monitorEvent();

        isEvent = false;
    } else if (isReadDB) {
        textPythonLower += textLine;

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
        textPythonLower += textLine;

        isWriteDB = false;
    } else if (!isEmptyLine) {
        if (isMPC) {
            textMpc += tab(1) + removeSpace(textLine);
        } else {
            textPythonLower += textLine;
        }
    }

    args.clear();
    isEmptyLine = true;

//    printf("-\n");
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