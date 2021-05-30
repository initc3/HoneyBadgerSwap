%option noyywrap

%{
#include <string>
#include <vector>
#include <assert.h>

using namespace std;

enum Phase {MONITOR, READDB, MPCINPUT, MPC, MPCOUTPUT, WRITEDB};
int phase;

bool isEvent, isReadDB, isWriteDB, isMpcInput, isMpcOutput, isEmptyLine = true;

int braces;
int numInput;
int cntValueDB;

string fileName;
string funcName;
string mpcFile;
FILE *mpcFp, *pythonFp;
int cntMPC;
string textLine;
string textPython, textMpc;

bool isArg;
string arg;
vector<string> args;

string tab(int repetition) {
    string ret = "";
    for (int i = 0; i < repetition; i++) {
        ret += "    ";
    }
    return ret;
}

void initPython() {
    textPython += "import asyncio\n";
    textPython += "import os\n";
    textPython += "import subprocess\n";
    textPython += "from ratel.src.python.utils import location_sharefile, prog, players, threshold, mpcPort, blsPrime, sz, int_to_hex, hex_to_int, recover_input, fp\n";
    textPython += "\n";

    string pythonFile = "python/" + fileName + ".py";
    pythonFp = fopen(pythonFile.c_str(), "w");
}

void finalizePython() {
    fprintf(pythonFp, "%s", textPython.c_str());
    fclose(pythonFp);
}

void initMPC() {
    mpcFile = "mpc/" + fileName + funcName + to_string(++cntMPC) + ".mpc";
    mpcFp = fopen(mpcFile.c_str(), "w");


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
    textPython += "\n";
    textPython += "async def monitor" + funcName + "(web3, db, serverID, contract, confirmation, account):\n";
    textPython += "    blkNum = web3.eth.get_block_number()\n";
    textPython += "    while True:\n";
    textPython += "        await asyncio.sleep(5)\n";
    textPython += "        curBlkNum = web3.eth.get_block_number()\n";
    textPython += "        if curBlkNum - blkNum > confirmation:\n";
    textPython += "            eventFilter = contract.events." + funcName + ".createFilter(fromBlock=blkNum, toBlock=curBlkNum - confirmation)\n";
    textPython += "            logs = eventFilter.get_all_entries()\n";
    textPython += "            blkNum = curBlkNum - confirmation + 1\n";
    textPython += "            for log in logs:\n";
    textPython += "\n";

    int num = args.size();
    for (int i = 0; i < num; i++) {
        string arg = args[i];
        textPython += tab(4) + arg + " = log['args']['" + arg + "']\n";
    }
    textPython += "\n";
}

void mpcInput() {
    //python
    textPython += "\n";
    textPython += tab(4) + "file = location_sharefile(serverID)\n";
    textPython += tab(4) + "with open(file, \"wb\") as f:\n";
    textPython += tab(4) + "    f.write(\n";
    numInput = args.size();
    for (int i = 0; i < numInput; i++) {
        textPython += tab(6) + ((i == 0) ? "" : "+ ") + "int_to_hex(" + args[i] + ")\n";
    }
    textPython += tab(4) + "    )\n";
    textPython += "\n";

    //mpc
    initMPC();
    for (int i = 0; i < numInput; i++) {
        textMpc += tab(1) + args[i] + " = read_sint(" + to_string(i) + ")\n";
    }
    textMpc += "\n";

    //python
    textPython += tab(4) + "env = os.environ.copy()\n";
    textPython += tab(4) + "cmd = [prog, \'-N\', f\'{players}\', \'-T\', f\'{threshold}\', \'-p\', f\'{serverID}\', \'-pn\', f\'{mpcPort()}\', \'-P\', f\'{blsPrime}\', \'" + mpcFile.substr(4, mpcFile.length() - 8) + "\']\n";
    textPython += tab(4) + "task = subprocess.Popen(cmd, env=env)\n";
    textPython += tab(4) + "task.wait()\n";
    textPython += "\n";
}

void mpcOutput() {
    //mpc
    textMpc += "\n";
    int num = args.size();
    for (int i = 0; i < num; i++) {
        textMpc += tab(1) + "writesharestofile(" + args[i] + ")\n";
    }
    textMpc += "\n";

    finalizeMPC();

    //python
    textPython += tab(4) + "input_argeter_num = " + to_string(numInput) + "\n";
    textPython += tab(4) + "with open(file, \"rb\") as f:\n";
    textPython += tab(4) + "    f.seek(input_argeter_num * sz)\n";
    for (int i = 0; i < num; i++) {
        textPython += tab(5) + args[i] + " = hex_to_int(f.read(sz))\n";
    }
    textPython += "\n";
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
    phase = 0;
    cntMPC = 0;
    cntValueDB = 0;
}

^.*monitorEvent\( {
    isEvent = true;
    isArg = true;
}

^.*readDB\( {
    string text = string(yytext);

    isReadDB = true;
    isArg = true;

    textLine += tab(3) + text.substr(0, text.length() - 7);
}

^.*writeDB\( {
    string text = string(yytext);

    isWriteDB = true;
    isArg = true;

    textLine += text.substr(0, text.length() - 8);
}

^.*mpcInput\( {
    isMpcInput = true;
    isArg = true;
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
        int pos = 0;
        for (int i = 0; i < arg.length(); i++) {
            if (arg[i] == ' ') {
                pos = i + 1;
                break;
            }
        }
        args.push_back(arg.substr(pos));
        arg = "";

        if (isReadDB) {
            assert(args.size() == 1);
            
            string textTmp = textLine;

            string value = "value" + to_string(++cntValueDB);
            
            textLine = tab(4) + "try:\n";
            textLine += tab(4) + "    " + value + " = db.Get(" + args[0] + ".encode())\n";
            textLine += tab(4) + "except KeyError:\n";
            textLine += tab(4) + "    " + value + " = bytes(0)\n";
            textLine += textTmp + value;
        } else if (isWriteDB) {
            assert(args.size() == 2);
            textLine += "db.Put(" + args[0] + ".encode(), " + args[1] + ")";
        }

        isArg = false;
    } else {
        addText(yytext);
    }
}

,{SPACE}* {
    if (isArg) {
        int pos = 0;
        for (int i = 0; i < arg.length(); i++) {
            if (arg[i] == ' ') {
                pos = i + 1;
                break;
            }
        }
        args.push_back(arg.substr(pos));
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
    if (braces == 0) {
//        addText("\n");
    } else {
        addText(yytext);
    }
}

\n {
    textLine += "\n";

//    printf("phase:%d-textLine:%s", phase, textLine.c_str());

    if (isEvent) {
        assert(phase == MONITOR);
        phase++;

        monitorEvent();

        isEvent = false;
    } else if (isReadDB) {
        assert(phase == READDB);

        textPython += textLine;

        isReadDB = false;
    } else if (isMpcInput) {
        assert(phase == READDB);
        phase++;

        mpcInput();

        isMpcInput = false;
    } else if (isMpcOutput) {
        assert(phase == MPC);
        phase++;

        mpcOutput();

        isMpcOutput = false;
    } else if (isWriteDB) {
        assert(phase == MPCOUTPUT || phase == WRITEDB);
        if (phase == MPCOUTPUT) phase++;

        textPython += tab(3) + textLine;

        isWriteDB = false;
    } else if (!isEmptyLine) {
        if (phase == MPCINPUT) {
            phase++;
            textMpc += textLine;
        } else if (phase == MPC) {
            textMpc += textLine;
        } else {
            textPython += tab(3) + textLine;
        }
    }

    args.clear();
    isEmptyLine = true;

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

	return 0;
}