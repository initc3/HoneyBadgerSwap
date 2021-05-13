%{
#include <string>

using namespace std;

bool isDec, isFunc;
int braces;

string textSol;

void condOutput(string text) {
    if (isDec) {
        if (braces > 0) {
            textSol += text;
        }
    } else if (isFunc) {
        textSol += text;
    }
}

%}

SPACE [ \n\t\r]+

%%

^declaration{SPACE} {
    isDec = true;
}

^func{SPACE} {
    isFunc = true;
    textSol += "\tfunction ";
}

\{ {
    string text = string(yytext);

    if (isFunc) {
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
        textSol += "\n";
        isDec = false;
        isFunc = false;
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
    string fileName = string(_fileName);
    fileName[0] = toupper(fileName[0]);

    textSol += "pragma solidity ^0.5.0;\n\n";
    textSol += "contract " + fileName + " {";

	yylex();

    textSol += "}\n";

    FILE *solFp;
    solFp = fopen("contracts/tmp.sol", "w");
    fprintf(solFp, "%s", textSol.c_str());
    fclose(solFp);

	return 0;
}