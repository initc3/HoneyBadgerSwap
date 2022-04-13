%option noyywrap

%{


#include <string>
#include <vector>
#include <assert.h>
#include <set>

using namespace std;

int braces;

string textLine;

string textPublicEvents;
string textPublicExtraFunctions;
string textPublic;

string textPrivate;

string textRecover;

bool isArg;
string arg;
vector<string> args;

bool isFunc;
bool isFuncParam;
string funcName;

bool isMpc;
bool isMpcParam;
int mpcBracesPos;
int mpcIndentLen;

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


string getIndent(string st) {
    for (int i = 0; i < st.length(); i++) {
        if (st[i] != ' ') {
            return st.substr(0, i);
        }
    }
}


string getFuncName(string st) {
    int len = st.length();
    for (int i = len - 1; i >= 0; i--) {
        if (st[i] == ' ') {
            string name = st.substr(i + 1, len - i - 2);
            name[0] = toupper(name[0]);
            return name;
        }
    }
    return "";
}


void addEventDeclare() {
    textPublicEvents += "    event " + funcName + "(uint seq" + funcName + ", ";
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textPublicEvents += ", ";
        if (args[i][0] == '$') {
            string varType = getType(args[i]).substr(1);
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textPublicEvents += "uint idx" + postVarName + ", " + varType + " masked" + postVarName;
        } else {
            textPublicEvents += args[i];
        }
    }
    textPublicEvents += ");\n";
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
            textRecover += "uint256," + varType;
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

    textPublic += "        uint seq" + funcName + " = opCnt++;\n";
    textPublic += "        opEvent[seq" + funcName + "] = \"" + funcName + "\";\n";
    textPublic += "        opContent[seq" + funcName + "] = abi.encode(";
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textPublic += ", ";
        if (args[i][0] == '$') {
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textPublic += "idx" + postVarName + ", masked" + postVarName;
        } else {
            textPublic += removeType(args[i]);
        }
    }
    textPublic += ");\n";
    textPublic += "\n";

    textPublic += "        emit " + funcName + "(seq" + funcName + ", ";
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textPublic += ", ";
        if (args[i][0] == '$') {
            string postVarName = removeType(args[i]);
            postVarName[0] = toupper(postVarName[0]);
            textPublic += "idx" + postVarName + ", " + "masked" + postVarName;
        } else {
            textPublic += removeType(args[i]);
        }
    }
    textPublic += ");\n";
}


void addMpcParam() {
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textLine += ", ";
        textLine += args[i];
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


void addInputMaskCheck() {
    bool flag = false;
    for (int i = 0; i < args.size(); i++) {
        string varName = removeType(args[i]);
        varName[0] = toupper(varName[0]);
        if (args[i][0] == '$') {
            textLine += "        require(inputMaskOwner[idx" + varName + "] == msg.sender);\n";
            flag = true;
        }
    }
    if (flag) textLine += "\n";
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

        textPublicExtraFunctions += "    function " + varName + "Set(";
        for (int i = 1; i < args.size(); i++) {
            if (i > 1) textPublicExtraFunctions += ", ";
            textPublicExtraFunctions += args[i];
        }
        textPublicExtraFunctions += ") public {\n";
        textPublicExtraFunctions += "        address server = msg.sender;\n";
        textPublicExtraFunctions += "        require(isServer[server]);\n";
        if (type == "string memory") {
            textPublicExtraFunctions += "        if (bytes(" + varName + "Value[server]).length > 0) {\n";
            textPublicExtraFunctions += "            " + varName + "Count[" + varName + "Value[server]]--;\n";
            textPublicExtraFunctions += "        }\n";
        } else if (type == "uint") {
            textPublicExtraFunctions += "        if (" + varName + "Value[server] != 0) {\n";
            textPublicExtraFunctions += "            " + varName + "Count[" + varName + "Value[server]]--;\n";
            textPublicExtraFunctions += "        }\n";
        } else if (type == "address") {
            textPublicExtraFunctions += "        if (" + varName + "Value[server] != address(0)) {\n";
            textPublicExtraFunctions += "            " + varName + "Count[" + varName + "Value[server]]--;\n";
            textPublicExtraFunctions += "        }\n";
        }
        textPublicExtraFunctions += "        " + varName + "Value[server] = " + value + ";\n";
        textPublicExtraFunctions += "        " + varName + "Count[" + varName + "Value[server]]++;\n";
        textPublicExtraFunctions += "        if (" + varName + "Count[" + value + "] > T) {\n";
        textPublicExtraFunctions += "            " + varName + parameter + " = " + value + ";\n";
        textPublicExtraFunctions += "        }\n";
        textPublicExtraFunctions += "    }\n";
        textPublicExtraFunctions += "\n";
        textPublicExtraFunctions += "\n";
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

        textPublicExtraFunctions += "    function " + opName + "Add(";
        for (int i = 1; i < args.size(); i++) {
          if (i > 1) textPublicExtraFunctions += ", ";
          textPublicExtraFunctions += args[i];
        }
        textPublicExtraFunctions += ") public {\n";
        textPublicExtraFunctions += "        address server = msg.sender;\n";
        textPublicExtraFunctions += "        require(isServer[server]);\n";
        textPublicExtraFunctions += "        require(" + opName + "Value[server][" + removeType(args[1]) + "] == 0);\n";
        textPublicExtraFunctions += "        require(" + opName + "Finish[" + removeType(args[1]) + "] == false);\n";
        textPublicExtraFunctions += "        " + opName + "Value[server][" + removeType(args[1]) + "] = " + value + ";\n";
        textPublicExtraFunctions += "        " + opName + "Count[" + removeType(args[1]) + "][" + value + "]++;\n";
        textPublicExtraFunctions += "        if (" + opName + "Count[" + removeType(args[1]) + "][" + value + "] > T) {\n";
        textPublicExtraFunctions += "            " + opName + "Finish[" + removeType(args[1]) + "] = true;\n";
        textPublicExtraFunctions += "            " + varName + parameter + " += " + value + ";\n";
        textPublicExtraFunctions += "        }\n";
        textPublicExtraFunctions += "    }\n";
        textPublicExtraFunctions += "\n";
        textPublicExtraFunctions += "\n";
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
    textPrivate += indent + "tx = server.contract.functions." + opName + op + "(";
    for (int i = 1; i < args.size(); i++) {
        if (i > 1) textPrivate += ", ";
        textPrivate += removeType(args[i]);
    }
    textPrivate += ").buildTransaction({'from': server.account.address, 'gas': 1000000, 'nonce': server.web3.eth.get_transaction_count(server.account.address)})\n";
    textPrivate += indent + "signedTx = server.web3.eth.account.sign_transaction(tx, private_key=server.account.privateKey)\n";
    textPrivate += indent + "server.web3.eth.send_raw_transaction(signedTx.rawTransaction)\n";
    textPrivate += indent + "server.web3.eth.wait_for_transaction_receipt(signedTx.hash)\n";
    textPrivate += indent + "print(server.contract.functions." + varName + "(";

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

    funcName = getFuncName(string(yytext));

    textLine += string(yytext);
}


^.*mpc\( {
    isMpc = true;
    isMpcParam = true;
    isArg = true;
    mpcBracesPos = braces;

    string text = string(yytext);
    mpcIndentLen = getIndent(text).length();
    textLine += text.substr(0, text.length() - 1) + " " + funcName + "(";
}


^.*pureMpc.*\( {
    isMpc = true;
    mpcBracesPos = braces;

    funcName = getFuncName(string(yytext));

    string text = string(yytext);
    mpcIndentLen = getIndent(text).length();
    textLine += text.substr(0, text.length() - 1 - funcName.length()) + funcName + "(";
}


^.*set\( {
    isSetData = true;
    isArg = true;

    indent = getIndent(string(yytext)).substr(mpcIndentLen);
}


^.*add\( {
    isAddData = true;
    isArg = true;

    indent = getIndent(string(yytext)).substr(mpcIndentLen);
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
        textPublic += "\n" + textPublicEvents + "\n\n" + textPublicExtraFunctions;
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

    //printf("textLine:%s", textLine.c_str());

    if (isMpc) {
        if (braces == mpcBracesPos) {
            isMpc = false;
            textLine += "\n";
        }

        if (isMpcParam) {
            addEventDeclare();
            addEventEmit();

            isMpcParam = false;
            args.clear();
            textPrivate += textLine.substr(mpcIndentLen);
        } else if (isSetData) {
            setData();
            callFunc("Set");

            isSetData = false;
            args.clear();
        } else if (isAddData) {
            addData();
            callFunc("Add");

            isAddData = false;
            args.clear();
        } else {
            textPrivate += (textLine.length() > mpcIndentLen ? textLine.substr(mpcIndentLen): textLine);
        }
    } else {
        if (isFunc) {
            if (braces == 1) {
                isFunc = false;
            }

            if (isFuncParam) {
                addInputMaskCheck();

                isFuncParam = false;
                args.clear();
            }
        }
        textPublic += textLine;
    }

    //printf("-\n");

    textLine = "";
}
%%


int main(int argc, char **argv) {
    char *_fileName = argv[1];
    string fileName = string(_fileName);

    textRecover += "from eth_abi import decode_single\n";
    textRecover += "\n";
    textRecover += "def recover(contract, opSeq):\n";
    textRecover += "    opEvent = contract.functions.opEvent(opSeq).call()\n";
    textRecover += "    opContent = contract.functions.opContent(opSeq).call()\n";
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
