%option noyywrap

%{
#include <string>

using namespace std;

string fileName;

bool isImport, isDec, isFunc;
int braces;

bool isImportFile;
string importFile;

string textSol;

void condOutput(string text) {
    if (isImport) {
        if (isImportFile) {
            importFile += text;
        }
    } else if (isFunc || braces > 0) {
        textSol += text;
    }
}

%}

SPACE [ \n\t\r]+

%%

^import{SPACE} {
    isImport = true;
}

^declaration{SPACE} {
    isDec = true;
}

^func{SPACE} {
    isFunc = true;
    textSol += "\tfunction ";
}

\{ {
    string text = string(yytext);

    if (isFunc && braces == 0) {
        textSol += "public ";
    }
    condOutput(text);
    braces++;
}

\} {
    string text = string(yytext);

    braces--;
    condOutput(text);
    if (braces == 0) {
        if (isImport) {
            textSol += "\ncontract " + fileName + " {";
            isImport = false;
        }
        else {
            textSol += "\n\n";
            isDec = false;
            isFunc = false;
        }
    }
}

\" {
    if (isImport) {
        if (isImportFile) {
            textSol += "import \"" + importFile + "\";\n";
            importFile = "";
        }
        isImportFile = !isImportFile;
    }
}

\n {
    string text = string(yytext);

    condOutput(text);
    if (isFunc) {
        textSol += "\t";
    }
}

. {
    string text = string(yytext);

    condOutput(yytext);
}
%%

int main(int argc, char **argv) {
    char *_fileName = argv[1];
    fileName = string(_fileName);
    fileName[0] = toupper(fileName[0]);

    textSol += "pragma solidity ^0.5.0;\n\n";

	yylex();

    textSol += "}\n";

    FILE *solFp;
    solFp = fopen("contracts/tmp.sol", "w");
    fprintf(solFp, "%s", textSol.c_str());
    fclose(solFp);

	return 0;
}