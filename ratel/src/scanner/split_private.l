%{
#include <set>
#include <string>
#include <vector>

using namespace std;

const string TAB = "    ";

enum Phase {READDB, MPCINPUT, MPC, MPCOUTPUT, WRITEDB};
int phase;

bool isFuncName, isFuncDec;
string funcName;
int braces;

bool isReadDB, isWriteDB;
bool isKey, isValue;
string key, value;

string textPython, textMpc;
string textLine, textBlock;
bool isEmptyLine = true;

bool isParam;
string param;

int numInput;
bool isMpcInput, isMpcOutput;
vector<string> mpcParams;

set<string> params;

void setData() {
    textPython += "def " + funcName + "SetData" + "(serverID";
    for (set<string>::iterator it = params.begin(); it != params.end(); ++it) {
        textPython += ", " + *it;
    }
    textPython += "):\n";
    textPython += TAB + "db = openDB(location_db(serverID))\n\n";
    textPython += textBlock + "\n";

    textBlock = "";
    params.clear();
}

void orgData() {
    textPython += "def " + funcName + "OrgData" + "(serverID";
    for (set<string>::iterator it = params.begin(); it != params.end(); ++it) {
        textPython += ", " + *it;
    }
    textPython += "):\n";
    textPython += TAB + "db = openDB(location_db(serverID))\n\n";
    textPython += textBlock + "\n";

    textBlock = "";
    params.clear();
}

void mpc() {
    textMpc += "def main():\n";
    textMpc += textBlock;
    textMpc += "\nmain()\n";

    textBlock = "";
}

%}

SPACE [ \t\r]+

%%

^func{SPACE} {
    isFuncName = true;
    isFuncDec = true;
    funcName = "";
}

^.*readDB\( {
    string text = string(yytext);

    isReadDB = true;
    isKey = true;

    textLine += text.substr(0, text.length() - 7);
}

^.*writeDB\( {
    string text = string(yytext);

    isWriteDB = true;
    isKey = true;

    textLine += text.substr(0, text.length() - 8);
}

^.*mpcInput\( {
    isMpcInput = true;
    isParam = true;
}

^.*mpcOutput\( {
    isMpcOutput = true;
    isParam = true;
}

\( {
    if (isFuncName) {
        isFuncName = false;
        isParam = true;
    }
}

\) {
    if (isReadDB) {
        textLine += "db.Get(" + key + ".encode())";

        isKey = false;
        key = "";
    } else if (isWriteDB) {
        textLine += "db.Put(" + key + ".encode(), " + value + ")";

        isValue = false;
        key = "";
        value = "";
    } else if (isMpcInput) {
        mpcParams.push_back(param);
        isParam = false;
        param = "";
    } else if (isMpcOutput) {
        mpcParams.push_back(param);
        isParam = false;
        param = "";
    } else if (isFuncDec && isParam) {
        for (int i = 0; i < param.length(); i++) {
            if (param[i] == ' ') {
                param = param.substr(i + 1);
                break;
            }
        }
        params.insert(param);
        param = "";
        isParam = false;
    }
}

,{SPACE}* {
    if (isKey) {
        isKey = false;
        isValue = true;
    } else if (isMpcInput || isMpcOutput) {
        mpcParams.push_back(param);
        param = "";
    } else if (isFuncDec && isParam) {
        for (int i = 0; i < param.length(); i++) {
            if (param[i] == ' ') {
                param = param.substr(i + 1);
                break;
            }
        }
        params.insert(param);
        param = "";
    }
}

\{ {
    braces++;
    if (isKey) {
        key += "{";
        isParam = true;
    }
}

\} {
    braces--;
    if (braces == 0) {
        assert(phase == 4);

        orgData();
    }
    if (isKey) {
        key += "}";
        if (isParam) {
            params.insert(param);
            isParam = false;
            param = "";
        }
    }
}

. {
    string text = string(yytext);

    if (yytext[0] != ' ' && yytext[0] != '\r' && yytext[0] != '\t') {
        isEmptyLine = false;
    }

    bool flag = true;
    if (isFuncName) {
        funcName += text;
        flag = false;
    }
    if (isKey) {
        key += text;
        flag = false;
    }
    if (isValue) {
        value += text;
        flag = false;
    }
    if (isParam) {
        param += text;
        flag = false;
    }

    if (flag) {
        textLine += text;
    }
}

\n {
    textLine += "\n";

    if (isReadDB) {
        assert(phase == 0);

        textBlock += textLine;

        isReadDB = false;
    } else if (isMpcInput) {
        assert(phase == 0);
        phase++;

        textLine += TAB + "file = location_sharefile(serverID)\n";
        textLine += TAB + "with open(file, \"wb\") as f:\n";
        textLine += TAB + TAB + "f.write(\n";
        int num = mpcParams.size();
        numInput = num;
        for (int i = 0; i < num; i++) {
            textLine += TAB + TAB + TAB + ((i == 0) ? "" : "+ ") + mpcParams[i] + "\n";
        }
        textLine += TAB + TAB + ")\n";
        textBlock += textLine;

        setData();

        for (int i = 0; i < num; i++) {
            textBlock += TAB + mpcParams[i] + " = read_sint(" + to_string(i) + ")\n";
        }
        textBlock += "\n";

        mpcParams.clear();
        isMpcInput = false;
    } else if (isMpcOutput) {
        assert(phase == 2);
        phase++;

        int num = mpcParams.size();
        for (int i = 0; i < num; i++) {
            textLine += TAB + "writesharestofile(" + mpcParams[i] + ")\n";
        }
        textBlock += textLine;

        mpc();

        textBlock += TAB + "file = location_sharefile(serverID)\n";
        textBlock += TAB + "input_parameter_num = " + to_string(numInput) + "\n";
        textBlock += TAB + "with open(file, \"rb\") as f:\n";
        textBlock += TAB + TAB + "f.seek(input_parameter_num * sz)\n";
        for (int i = 0; i < num; i++) {
            textBlock += TAB + TAB + mpcParams[i] + " = f.read(sz)\n";
        }
        textBlock += "\n";

        mpcParams.clear();
        isMpcOutput = false;
    } else if (isWriteDB) {
        assert(phase == 3 || phase == 4);
        if (phase == 3) phase++;

        textBlock += textLine;

        isWriteDB = false;
    } else if (!isEmptyLine && !isFuncDec) {
        assert(phase == 1 || phase == 2);
        if (phase == 1) phase++;

        textBlock += textLine;
    }
    isEmptyLine = true;
    isFuncDec = false;

//    printf("textLine:%sphase:%d\n", textLine.c_str(), phase);

    textLine = "";
}
%%

int main(int argc, char **argv) {
    char *_fileName = argv[1];
    string fileName = string(_fileName);

    textPython += "from .utils import openDB, location_db, location_sharefile, sz\n\n";

    textMpc += "def read_sint(idx):\n" +
               TAB + "start = regint(idx * 32)\n" +
               TAB + "end = regint()\n" +
               TAB + "share = sint()\n" +
               TAB + "readsharesfromfile(start, end, share)\n" +
               TAB + "return share\n\n";

    yylex();

    FILE *pythonFp;
    string pythonFile = "python/" + fileName + ".py";
    pythonFp = fopen(pythonFile.c_str(), "w");
//    printf("%s", textPython.c_str());
    fprintf(pythonFp, "%s", textPython.c_str());
    fclose(pythonFp);

    FILE *mpcFp;
    mpcFp = fopen("mpc/tmp.mpc", "w");
//    printf("%s", textMpc.c_str());
    fprintf(mpcFp, "%s", textMpc.c_str());
    fclose(mpcFp);

	return 0;
}