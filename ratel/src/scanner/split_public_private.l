%option noyywrap

%{
#include <string>
#include <vector>
#include <assert.h>

using namespace std;

int braces;

string textLine, textPublic, textPublicFunc, textPrivate;

bool isFunc;
bool isFuncName;
string funcName;
bool isFuncParam;

bool isArg;
string arg;
vector<string> args;

bool isMpc;
int pos;
bool isMpcHead;

bool isSetData;

string removeType(string st) {
    for (int i = st.length() - 1; i >= 0; i--) {
        if (st[i] == ' ') {
            return st.substr(i + 1);
        }
    }
}

%}

SPACE [ \t\r]+

%%

^.*function{SPACE} {
    isFunc = true;
    isFuncName = true;
    isArg = true;

    textLine += string(yytext);
}

^.*mpc\( {
    isMpc = true;
    isMpcHead = true;
    pos = braces;
    isArg = true;

    string text = string(yytext);
    textLine += text.substr(0, text.length() - 1) + " " + funcName + "(";
}

^.*set\( {
    isSetData = true;
    isArg = true;
}

,{SPACE}* {
    if (isArg) {
        args.push_back(arg);
        arg = "";
    } else {
        textLine += string(yytext);
    }
}

\( {
    if (isFuncName) {
        funcName = arg;
        textLine += funcName;
        funcName[0] = toupper(funcName[0]);
        isFuncName = false;

        arg = "";
        isFuncParam = true;
    }
    textLine += string(yytext);
}

\) {
    if (isArg) {
        args.push_back(arg);
        arg = "";
        isArg = false;

        if (isMpc) {
            if (!isSetData) {
                textPublic += "    event " + funcName + "(";
                textPublicFunc += "        emit " + funcName + "(";
                for (int i = 0; i < args.size(); i++) {
                    if (i > 0) {
                        textPublic += ", ";
                        textPublicFunc += ", ";
                        textLine += ", ";
                    }
                    if (args[i][0] == '$') {
                        string varType, varName;
                        for (int j = 0; j < args[i].length(); j++) {
                            if (args[i][j] == ' ') {
                                varType = args[i].substr(1, j - 1);
                                varName = args[i].substr(j + 1);
                                varName[0] = toupper(varName[0]);
                                break;
                            }
                        }
                        textPublic += "uint idx" + varName + ", " + varType + " masked" + varName;
                        textPublicFunc += "idx" + varName + ", " + "masked" + varName;
                        textLine += "idx" + varName + ", " + "masked" + varName;
                    } else {
                        textPublic += args[i];
                        textPublicFunc += removeType(args[i]);
                        textLine += removeType(args[i]);
                    }
                }
                textPublic += ");\n";
                textPublicFunc += ");\n";
            }
        } else {
            assert(isFuncParam);
            for (int i = 0; i < args.size(); i++) {
                if (i > 0) {
                    textLine += ", ";
                }
                if (args[i][0] == '$') {
                    string varType, varName;
                    for (int j = 0; j < args[i].length(); j++) {
                        if (args[i][j] == ' ') {
                            varType = args[i].substr(1, j - 1);
                            varName = args[i].substr(j + 1);
                            varName[0] = toupper(varName[0]);
                            break;
                        }
                    }
                    textLine += "uint idx" + varName + ", " + varType + " masked" + varName;
                } else {
                    textLine += args[i];
                }
            }
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
        textPublic += "\n";
        textPublic += textPublicFunc;
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

//    printf("textLine:%s\n", textLine.c_str());

    if (isFunc) {
        if (isMpc) {
            if (isMpcHead) {
                for (int i = 0; i < args.size(); i++) {
                    if (args[i][0] == '$') {
                        string varName = removeType(args[i]);
                        string postVarName = varName;
                        postVarName[0] = toupper(postVarName[0]);
                        textLine += "    " + varName + " = recover_input(db, masked" + postVarName + ", idx" + postVarName + ")\n";
                    }
                }

                isMpcHead = false;
                args.clear();
                textPrivate += textLine.substr(8);
            } else if (isSetData) {
                string varName = args[0];

                string tmp = varName;
                tmp[0] = toupper(tmp[0]);
                string valueVarName = "value" + tmp;
                string countVarName = "count" + tmp;

                string value = removeType(args[1]);
                string parameter = "";
                for (int i = 2; i < args.size(); i++) {
                    parameter += "[" + removeType(args[i]) + "]";
                }

                string textTmp;
                textTmp += "\n";
                textTmp += "    function set" + tmp + "(";
                for (int i = 1; i < args.size(); i++) {
                    if (i > 1) {
                        textTmp += ", ";
                    }
                    textTmp += args[i];
                }
                textTmp += ") public {\n";

                textTmp += "        address server = msg.sender;\n";
                textTmp += "        require(isServer[server]);\n";
                textTmp += "        if (bytes(" + valueVarName + "[server]).length > 0) {\n";
                textTmp += "            " + countVarName + "[" + valueVarName + "[server]]--;\n";
                textTmp += "        }\n";
                textTmp += "        " + valueVarName + "[server] = " + value + ";\n";
                textTmp += "        " + countVarName + "[" + valueVarName + "[server]]++;\n";
                textTmp += "        if (" + countVarName + "[" + value + "] > T) {\n";
                textTmp += "            " + varName + parameter + " = " + value + ";\n";
                textTmp += "        }\n";
                textTmp += "    }\n";
                textTmp += textPublicFunc;
                textPublicFunc = textTmp;

                textPrivate += "\n";
                textPrivate += "    tx = contract.functions.set" + tmp + "(";
                for (int i = 1; i < args.size(); i++) {
                    if (i > 1) {
                        textPrivate += ", ";
                    }
                    textPrivate += removeType(args[i]);
                }
                textPrivate += ").buildTransaction({'from': account.address, 'gas': 1000000, 'nonce': web3.eth.get_transaction_count(account.address)})\n";
                textPrivate += "    signedTx = web3.eth.account.sign_transaction(tx, private_key=account.privateKey)\n";
                textPrivate += "    web3.eth.send_raw_transaction(signedTx.rawTransaction)\n";
                textPrivate += "    web3.eth.wait_for_transaction_receipt(signedTx.hash)\n";
                textPrivate += "    print(contract.functions." + varName + "(";

                for (int i = 2; i < args.size(); i++) {
                    if (i > 2) {
                        textPrivate += ", ";
                    }
                    textPrivate += removeType(args[i]);
                }
                textPrivate += ").call())\n";

                isSetData = false;
                args.clear();
            } else textPrivate += (textLine.length() > 8 ? textLine.substr(8): textLine);
            if (braces == pos) {
                isMpc = false;
            }
        } else {
            textPublicFunc += textLine;
            if (braces == 1) {
                isFunc = false;
                textPublicFunc += "\n";
            }
            if (isFuncParam) {
                for (int i = 0; i < args.size(); i++) {
                    string varName = removeType(args[i]);
                    varName[0] = toupper(varName[0]);
                    if (args[i][0] == '$') {
                        textPublicFunc += "        require(inputMaskOwner[idx" + varName + "] == msg.sender);\n";
                    }
                }
                textPublicFunc += "\n";
                isFuncParam = false;
                args.clear();
            }
        }
    } else {
        textPublic += textLine;
    }

    textLine = "";
}
%%

int main(int argc, char **argv) {
    char *_fileName = argv[1];
    string fileName = string(_fileName);

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

	return 0;
}