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
    for (int i = 0; i < st.length(); i++) {
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


string argWithType(string arg, bool isEvent=false) {
    string ret;
    if (arg[0] == '$') {
        string varType = getType(arg).substr(1);
        string postVarName = removeType(arg);
        postVarName[0] = toupper(postVarName[0]);
        string tmp = "";
        string mem = isEvent ? "" : "memory" ;
        if (varType[0] == '#') {
            varType = varType.substr(1);
            tmp = ", uint idx" + postVarName + "Blinding, uint masked" + postVarName + "Blinding, uint8[] " + mem + " proof" + postVarName + ", uint8[] " + mem + " commitment" + postVarName;
        }
        ret = "uint idx" + postVarName + ", uint masked" + postVarName + tmp;
    } else {
        ret = arg;
    }
    return ret;
}


string argWithoutType(string arg) {
    string ret;
    if (arg[0] == '$') {
        string postVarName = removeType(arg);
        postVarName[0] = toupper(postVarName[0]);
        ret = "idx" + postVarName + ", masked" + postVarName;
        if (arg[1] == '#') {
            ret += ", idx" + postVarName + "Blinding, masked" + postVarName + "Blinding, proof" + postVarName + ", commitment" + postVarName;
        }
    } else {
        ret = removeType(arg);
    }
    return ret;
}


void addEventDeclare() {
    textPublicEvents += "    event " + funcName + "(uint seq" + funcName;
    for (int i = 0; i < args.size(); i++) {
        textPublicEvents += ", " + argWithType(args[i], true);

    }
    textPublicEvents += ");\n";
}


void addEventEmit() {
    textRecover += "    if opEvent == \'" + funcName + "\':\n";
    textRecover += "        ";
    if (!args.size()) {
        textRecover += "pass\n\n";
    } else {
        for (int i = 0; i < args.size(); i++) {
            textRecover += argWithoutType(args[i]) + ", ";
        }
        textRecover += " = decode_single(\'(";
        for (int i = 0; i < args.size(); i++) {
            if (i > 0) textRecover += ",";
            textRecover += argWithType(args[i], true);
        }
        textRecover += ")\', opContent)\n";
        textRecover += "        return eval(f'\{op\}" + funcName + "')(opSeq";
        for (int i = 0; i < args.size(); i++) {
            textRecover += ", " + argWithoutType(args[i]);
        }
        textRecover += ")\n";
        textRecover += "\n";
    }

    textPublic += "        uint seq" + funcName + " = opCnt++;\n";
    textPublic += "        opEvent[seq" + funcName + "] = \"" + funcName + "\";\n";
    textPublic += "        opContent[seq" + funcName + "] = abi.encode(";
    for (int i = 0; i < args.size(); i++) {
        if (i > 0) textPublic += ", ";
        textPublic += argWithoutType(args[i]);
    }
    textPublic += ");\n";
    textPublic += "\n";

    textPublic += "        emit " + funcName + "(seq" + funcName ;
    for (int i = 0; i < args.size(); i++) {
        textPublic += ", " + argWithoutType(args[i]);
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
        textLine += argWithType(args[i]);
    }
}


void addInputMaskCheck() {
    bool flag = false;
    for (int i = 0; i < args.size(); i++) {
        string varName = removeType(args[i]);
        varName[0] = toupper(varName[0]);
        if (args[i][0] == '$') {
            textLine += "        require(inputMaskOwner[idx" + varName + "] == msg.sender);\n";
            if (args[i][1] == '#') {
                textLine += "        require(inputMaskOwner[idx" + varName + "Blinding] == msg.sender);\n";
            }
            flag = true;
        }
    }
    if (flag) textLine += "\n";
}


void declareVariable(string varName, string type, vector<string> args) {
    if (args.size() > 2) {
        textPublicExtraFunctions += "    ";
        for (int i = 2; i < args.size(); i++) {
            textPublicExtraFunctions += "mapping(" + getType(args[i]) + " => ";
        }
        textPublicExtraFunctions += type;
        for (int i = 2; i < args.size(); i++) {
            textPublicExtraFunctions += ")";
        }
        textPublicExtraFunctions += " public " + varName + ";\n";
    } else {
        textPublicExtraFunctions += "    " + type + " public " + varName + ";\n";
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

        string varNameValue = varName + "Value";
        string varNameCount = varName + "Count";

        /* declare variables */
        args.push_back("address ");
        declareVariable(varNameValue, type, args);
        args.pop_back();
        args.push_back(type + " ");
        declareVariable(varNameCount, "uint", args);
        args.pop_back();

        /* define function call */
        textPublicExtraFunctions += "    function " + varName + "Set(";
        for (int i = 1; i < args.size(); i++) {
            if (i > 1) textPublicExtraFunctions += ", ";
            textPublicExtraFunctions += args[i];
        }
        textPublicExtraFunctions += ") public onlyServer {\n";
        textPublicExtraFunctions += "        address server = msg.sender;\n";
        if (type == "string memory") {
            textPublicExtraFunctions += "        if (bytes(" + varNameValue + parameter + "[server]).length > 0) {\n";
            textPublicExtraFunctions += "            " + varNameCount + parameter + "[" + varNameValue + parameter + "[server]]--;\n";
            textPublicExtraFunctions += "        }\n";
        } else if (type == "uint") {
            textPublicExtraFunctions += "        if (" + varNameValue + parameter + "[server] != 0) {\n";
            textPublicExtraFunctions += "            " + varNameCount + parameter + "[" + varNameValue + parameter + "[server]]--;\n";
            textPublicExtraFunctions += "        }\n";
        } else if (type == "address") {
            textPublicExtraFunctions += "        if (" + varNameValue + parameter + "[server] != address(0)) {\n";
            textPublicExtraFunctions += "            " + varNameCount + parameter + "[" + varNameValue + parameter + "[server]]--;\n";
            textPublicExtraFunctions += "        }\n";
        }
        textPublicExtraFunctions += "        " + varNameValue + parameter + "[server] = " + value + ";\n";
        textPublicExtraFunctions += "        " + varNameCount + parameter + "[" + varNameValue + parameter + "[server]]++;\n";
        textPublicExtraFunctions += "        if (" + varNameCount + parameter + "[" + value + "] > T) {\n";
        textPublicExtraFunctions += "            " + varName + parameter + " = " + value + ";\n";
        textPublicExtraFunctions += "        }\n";
        textPublicExtraFunctions += "    }\n";
        textPublicExtraFunctions += "\n";
        textPublicExtraFunctions += "\n";
    }
}


void addData() {
    string varName = args[0];

    if (addDataPool.find(varName) == addDataPool.end()) {
        assert(args.size() > 2);
        addDataPool.insert(varName);

        string value = removeType(args[1]);
        string type = getType(args[1]);
        string parameterWithoutSeq = "";
        for (int i = 2; i < args.size() - 1; i++) {
            parameterWithoutSeq += "[" + removeType(args[i]) + "]";
        }
        string parameter = parameterWithoutSeq + "[" + removeType(args.back()) + "]";

        string varNameValue = varName + "Value";
        string varNameCount = varName + "Count";
        string varNameFinish = varName + "Finish";

        /* declare variables */
        args.push_back("address ");
        declareVariable(varNameValue, type, args);
        args.pop_back();
        args.push_back(type + " ");
        declareVariable(varNameCount, "uint", args);
        args.pop_back();
        declareVariable(varNameFinish, "bool", args);

        /* define function call */

        textPublicExtraFunctions += "    function " + varName + "Add(";
        for (int i = 1; i < args.size(); i++) {
          if (i > 1) textPublicExtraFunctions += ", ";
          textPublicExtraFunctions += args[i];
        }
        textPublicExtraFunctions += ") public onlyServer {\n";
        textPublicExtraFunctions += "        address server = msg.sender;\n";
        textPublicExtraFunctions += "        require(" + varNameValue + parameter + "[server] == 0);\n";
        textPublicExtraFunctions += "        require(" + varNameFinish + parameter + " == false);\n";
        textPublicExtraFunctions += "        " + varNameValue + parameter + "[server] = " + value + ";\n";
        textPublicExtraFunctions += "        " + varNameCount + parameter + "[" + value + "]++;\n";
        textPublicExtraFunctions += "        if (" + varNameCount + parameter + "[" + value + "] > T) {\n";
        textPublicExtraFunctions += "            " + varNameFinish + parameter + " = true;\n";
        textPublicExtraFunctions += "            " + varName + parameterWithoutSeq + " += " + value + ";\n";
        textPublicExtraFunctions += "        }\n";
        textPublicExtraFunctions += "    }\n";
        textPublicExtraFunctions += "\n";
        textPublicExtraFunctions += "\n";
    }
}


void callFunc(string op, int init) {
    string varName = args[0];

    textPrivate += "\n";
    textPrivate += indent + "tx = server.contract.functions." + varName + op + "(";
    for (int i = 1; i < args.size(); i++) {
        if (i > 1) textPrivate += ", ";
        textPrivate += removeType(args[i]);
    }
    textPrivate += ").buildTransaction({'from': server.account.address, 'gas': 1000000, 'nonce': server.web3.eth.get_transaction_count(server.account.address)})\n";
    textPrivate += indent + "sign_and_send(tx, server.web3, server.account)\n";
    textPrivate += indent + "print(server.contract.functions." + varName + "(";

    for (int i = 2; i < args.size() - init; i++) {
        if (i > 2) textPrivate += ", ";
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

//    printf("textLine:%s", textLine.c_str());

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
            callFunc("Set", 0);

            isSetData = false;
            args.clear();
        } else if (isAddData) {
            addData();
            callFunc("Add", 1);

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

    textLine = "";
}
%%


int main(int argc, char **argv) {
    char *_fileName = argv[1];
    string fileName = string(_fileName);

    textRecover += "from eth_abi import decode_single\n";
    textRecover += "\n";
    textRecover += "def recover(contract, opSeq, op):\n";
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
