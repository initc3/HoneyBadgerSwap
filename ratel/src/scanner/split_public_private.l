%option noyywrap

%{
#include <string>
#include <vector>
#include <assert.h>
#include <set>

using namespace std;

int braces;

string textLine, textPublic, textPublicFuncPre, textPublicFuncSuf, textPrivate;
string textRecover;

bool isArg;
string arg;
vector<string> args;

bool isConstructor;

bool isFunc;
string funcName;
bool isFuncParam;

bool isMpc;
int pos;
bool isMpcParam;

string indent;
bool isSetData, isAddData;
set<string> setDataPool, addDataPool;

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

void getIndent(string st) {
    for (int i = 0; i < st.length(); i++) {
        if (st[i] != ' ') {
            indent = st.substr(0, i - 12);
            return;
        }
    }
}

void getFuncName(string st) {
    int len = st.length();
    for (int i = len - 1; i >= 0; i--) {
        if (st[i] == ' ') {
            funcName = st.substr(i + 1, len - i - 2);
            funcName[0] = toupper(funcName[0]);
            return;
        }
    }
}

void addEventDeclare() {
    textPublic += "    event " + funcName + "(uint seq" + funcName + ", ";
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textPublic += ", ";
        if (args[i][0] == '$') {
            string varType = getType(args[i]).substr(1);
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textPublic += "uint idx" + postVarName + ", " + varType + " masked" + postVarName;
        } else {
            textPublic += args[i];
        }
    }
    textPublic += ");\n";
}

void addEventEmit() {
    textRecover += "    if opEvent == \'" + funcName + "\':\n";
    textRecover += "        ";
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textRecover += ", ";
        if (args[i][0] == '$') {
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textRecover += "idx" + postVarName + ", masked" + postVarName;
        } else {
            textRecover += removeType(args[i]);
        }
    }
    textRecover += " = decode_single(\'(";
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textRecover += ",";
        if (args[i][0] == '$') {
            string varType = getType(args[i]).substr(1);
            textRecover += "uint256, " + varType;
        } else {
            textRecover += getType(args[i]);
        }
    }
    textRecover += ")\', opContent)\n";
    textRecover += "        return recover" + funcName + "(opSeq";
    for (int i = 0; i < args.size(); i++) {
        textRecover += ", ";
        if (args[i][0] == '$') {
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textRecover += "idx" + postVarName + ", masked" + postVarName;
        } else {
            textRecover += removeType(args[i]);
        }
    }
    textRecover += ")\n";
    textRecover += "\n";

    textPublicFuncSuf += "        uint seq" + funcName + " = opCnt++;\n";
    textPublicFuncSuf += "        opEvent[seq" + funcName + "] = \"" + funcName + "\";\n";
    textPublicFuncSuf += "        opContent[seq" + funcName + "] = abi.encode(";
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textPublicFuncSuf += ", ";
        if (args[i][0] == '$') {
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textPublicFuncSuf += "idx" + postVarName + ", masked" + postVarName;
        } else {
            textPublicFuncSuf += removeType(args[i]);
        }
    }
    textPublicFuncSuf += ");\n";

    textPublicFuncSuf += "        emit " + funcName + "(seq" + funcName + ", ";
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textPublicFuncSuf += ", ";
        if (args[i][0] == '$') {
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textPublicFuncSuf += "idx" + postVarName + ", " + "masked" + postVarName;
        } else {
            textPublicFuncSuf += removeType(args[i]);
        }
    }
    textPublicFuncSuf += ");\n";
}

void addMpcParam() {
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

void addFuncParam() {
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textLine += ", ";
        if (args[i][0] == '$') {
            string varType = getType(args[i]).substr(1);
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textLine += "uint idx" + postVarName + ", " + varType + " masked" + postVarName;
        } else {
            textLine += args[i];
        }
    }
}

void unmaskInput() {
    for (int i = 0; i < args.size(); i++) {
        if (args[i][0] == '$') {
            string varName = removeType(args[i]);
            string postVarName = varName;
            postVarName[0] = toupper(postVarName[0]);
            textLine += "    " + varName + " = recover_input(server.db, masked" + postVarName + ", idx" + postVarName + ")\n";
        }
    }
}

void setData() {
    string varName = args[0];

    if (setDataPool.find(varName) == setDataPool.end()) {
        setDataPool.insert(varName);

        string value = removeType(args[1]);
        string type = getType(args[1]);
        string parameter = "";
        for (int i = 2; i < args.size(); i++) {
            parameter += "[" + removeType(args[i]) + "]";
        }

        string textTmp = "    function " + varName + "Set(";
        for (int i = 1; i < args.size(); i++) {
            if (i > 1) textTmp += ", ";
            textTmp += args[i];
        }
        textTmp += ") public {\n";
        textTmp += "        address server = msg.sender;\n";
        textTmp += "        require(isServer[server]);\n";
        if (type == "string memory") {
            textTmp += "        if (bytes(" + varName + "Value[server]).length > 0) {\n";
            textTmp += "            " + varName + "Count[" + varName + "Value[server]]--;\n";
            textTmp += "        }\n";
        } else if (type == "uint") {
            textTmp += "        if (" + varName + "Value[server] != 0) {\n";
            textTmp += "            " + varName + "Count[" + varName + "Value[server]]--;\n";
            textTmp += "        }\n";
        } else if (type == "address") {
            textTmp += "        if (" + varName + "Value[server] != address(0)) {\n";
            textTmp += "            " + varName + "Count[" + varName + "Value[server]]--;\n";
            textTmp += "        }\n";
        }
        textTmp += "        " + varName + "Value[server] = " + value + ";\n";
        textTmp += "        " + varName + "Count[" + varName + "Value[server]]++;\n";
        textTmp += "        if (" + varName + "Count[" + value + "] > T) {\n";
        textTmp += "            " + varName + parameter + " = " + value + ";\n";
        textTmp += "        }\n";
        textTmp += "    }\n";
        textTmp += "\n";

        textTmp += textPublicFuncSuf;
        textPublicFuncSuf = textTmp;
    }
}

void addData() {
    string varName = args[0];
    string opName = removeType(args[1]);
    opName = opName.substr(3, opName.length());
    opName[0] = tolower(opName[0]);

    if (addDataPool.find(opName) == addDataPool.end()) {
        addDataPool.insert(opName);

        string value = removeType(args[2]);
        string parameter = "";
        for (int i = 3; i < args.size(); i++) {
          parameter += "[" + removeType(args[i]) + "]";
        }

        string textTmp = "    function " + opName + "Add(";
        for (int i = 1; i < args.size(); i++) {
          if (i > 1) textTmp += ", ";
          textTmp += args[i];
        }
        textTmp += ") public {\n";
        textTmp += "        address server = msg.sender;\n";
        textTmp += "        require(isServer[server]);\n";
        textTmp += "        require(" + opName + "Value[server][" + removeType(args[1]) + "] == 0);\n";
        textTmp += "        require(" + opName + "Finish[" + removeType(args[1]) + "] == false);\n";
        textTmp += "        " + opName + "Value[server][" + removeType(args[1]) + "] = " + value + ";\n";
        textTmp += "        " + opName + "Count[" + removeType(args[1]) + "][" + value + "]++;\n";
        textTmp += "        if (" + opName + "Count[" + removeType(args[1]) + "][" + value + "] > T) {\n";
        textTmp += "            " + opName + "Finish[" + removeType(args[1]) + "] = true;\n";
        textTmp += "            " + varName + parameter + " += " + value + ";\n";
        textTmp += "        }\n";
        textTmp += "    }\n";
        textTmp += "\n";

        textTmp += textPublicFuncSuf;
        textPublicFuncSuf = textTmp;
    }
}

void callFunc(string op) {
    string varName = args[0];
    string opName;
    int init;
    if (op == "Add") {
        opName = removeType(args[1]);
        opName = opName.substr(3, opName.length());
        opName[0] = tolower(opName[0]);
        init = 1;
    } else {
        assert(op == "Set");
        opName = args[0];
    }

    textPrivate += "\n";
    textPrivate += indent + "    tx = server.server.contract.functions." + opName + op + "(";
    for (int i = 1; i < args.size(); i++) {
        if (i > 1) textPrivate += ", ";
        textPrivate += removeType(args[i]);
    }
    textPrivate += ").buildTransaction({'from': server.account.address, 'gas': 1000000, 'nonce': server.web3.eth.get_transaction_count(server.account.address)})\n";
    textPrivate += indent + "    signedTx = server.web3.eth.account.sign_transaction(tx, private_key=server.account.privateKey)\n";
    textPrivate += indent + "    server.web3.eth.send_raw_transaction(signedTx.rawTransaction)\n";
    textPrivate += indent + "    server.web3.eth.wait_for_transaction_receipt(signedTx.hash)\n";
    textPrivate += indent + "    print(server.contract.functions." + varName + "(";

    for (int i = 2 + init; i < args.size(); i++) {
        if (i > 2 + init) textPrivate += ", ";
        textPrivate += removeType(args[i]);
    }
    textPrivate += ").call())\n";
    textPrivate += "\n";
}

%}

SPACE [ \t\r]+

%%

^.*function.*\( {
    isFunc = true;
    isFuncParam = true;
    isArg = true;

    getFuncName(string(yytext));

    textLine += string(yytext);
}

^.*constructor {
    isConstructor = true;
    textLine += string(yytext);
}

^.*mpc\( {
    isMpc = true;
    isMpcParam = true;
    pos = braces;
    isArg = true;

    string text = string(yytext);
    textLine += text.substr(0, text.length() - 1) + " " + funcName + "(";
}

^.*set\( {
    isSetData = true;
    isArg = true;

    getIndent(string(yytext));
}

^.*add\( {
    isAddData = true;
    isArg = true;

    getIndent(string(yytext));
}

,{SPACE}* {
    if (isArg) {
        args.push_back(arg);
        arg = "";
    } else {
        textLine += string(yytext);
    }
}

\) {
    if (isArg) {
        if (arg.length()) args.push_back(arg);
        arg = "";
        isArg = false;

        if (isMpcParam) {
            addMpcParam();
        }
        else if (isFuncParam) {
            addFuncParam();
        }
    }

    textLine += string(yytext);
}

\{ {
    braces++;

    textLine += string(yytext);
}

\} {
    braces--;

    if (braces == 0) {
        textPublic += textPublicFuncPre + textPublicFuncSuf;
    }
    textLine += string(yytext);
}

. {
    if (isArg) {
        arg += string(yytext);

    } else {
        textLine += string(yytext);
    }
}

\n {
    textLine += string(yytext);

//    printf("textLine:%s", textLine.c_str());

    if (isFunc) {
        if (isMpc) {
            if (braces == pos) {
                textPrivate += "    execHistory = readDB(f'execHistory', dict)\n";
                textPrivate += "    execHistory[seq" + funcName + "] = True\n";
                textPrivate += "    writeDB(f'execHistory', execHistory, dict)\n";
                isMpc = false;
            }
            if (isMpcParam) {
                addEventDeclare();
                addEventEmit();
                unmaskInput();

                isMpcParam = false;
                args.clear();
                textPrivate += textLine.substr(8);
            }
            else if (isSetData) {
                setData();
                callFunc("Set");

                isSetData = false;
                args.clear();
            }
            else if (isAddData) {
                addData();
                callFunc("Add");

                isAddData = false;
                args.clear();
            }
            else {
                textPrivate += (textLine.length() > 8 ? textLine.substr(8): textLine);
            }
        } else {
            textPublicFuncSuf += textLine;
            if (braces == 1) {
                isFunc = false;
                textPublicFuncSuf += "\n";
            }
            if (isFuncParam) {
                bool flag = false;
                for (int i = 0; i < args.size(); i++) {
                    string varName = removeType(args[i]);
                    varName[0] = toupper(varName[0]);
                    if (args[i][0] == '$') {
                        textPublicFuncSuf += "        require(inputMaskOwner[idx" + varName + "] == msg.sender);\n";
                        flag = true;
                    }
                }
                if (flag) textPublicFuncSuf += "\n";

                isFuncParam = false;
                args.clear();
            }
        }
    }
    else {
        if (isConstructor) {
            if (braces == 1) {
                isConstructor = false;
            }
            textPublicFuncPre += textLine;
        } else if (!(textLine.length() == 1 && textLine[0] == '\n')) textPublic += textLine;
    }

//    printf("-\n");
    textLine = "";
}
%%

int main(int argc, char **argv) {
    char *_fileName = argv[1];
    string fileName = string(_fileName);

    textRecover += "from eth_abi import decode_single\n";
    textRecover += "\n";
    textRecover += "def recover(contract, opSeq):\n";
    textRecover += "    opEvent = server.contract.functions.opEvent(opSeq).call()\n";
    textRecover += "    opContent = server.contract.functions.opContent(opSeq).call()\n";
    textRecover += "\n";

	yylex();

    FILE *solFp;
    string solFile = "contracts/" + fileName + ".sol";
    solFp = fopen(solFile.c_str(), "w");
    fprintf(solFp, "%s", textPublic.c_str());
    fclose(solFp);

    FILE *mpcFp;
    string mpcFile = "mpc/" + fileName + ".mpc";
	mpcFp = fopen(mpcFile.c_str(), "w");
	fprintf(mpcFp, "%s", textPrivate.c_str());
    fclose(mpcFp);

    FILE *recoverFp;
    string recoverFile = "python/" + fileName + "Recover.py";
    recoverFp = fopen(recoverFile.c_str(), "w");
	fprintf(recoverFp, "%s", textRecover.c_str());
    fclose(recoverFp);

	return 0;
}