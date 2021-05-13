%{
#include <string>

using namespace std;

bool isPublic, isPrivate;
int braces;

string textPublic, textPrivate;

void output(string text) {
    if (isPublic) {
        textPublic += text;
    } else if (isPrivate) {
        textPrivate += text;
    }
}

%}

SPACE [ \n\t\r]+

%%

^public{SPACE} {
    if (isPrivate) {
        printf("Something is wrong\n");
        exit(1);
    }
    isPublic = true;
}

^private{SPACE} {
    if (isPublic) {
        printf("Something is wrong\n");
        exit(1);
    }
    isPrivate = true;
}

\{ {
    string text = string(yytext);
    braces++;
    output(text);
}

\} {
    string text = string(yytext);
    output(text);
    braces--;
    if (braces == 0) {
        output("\n\n");
        isPublic = false;
        isPrivate = false;
    }
}

\n {
    string text = string(yytext);
    output(text);
}

. {
    string text = string(yytext);
    output(text);
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