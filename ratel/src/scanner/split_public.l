%option noyywrap

%{
#include <string>

using namespace std;

string textPublic;
bool otherParameters, isConstructorHead, isConstructorBody;

int braces;

%}

SPACE [ \t\r]*

%%

.*constructor{SPACE}\( {
    textPublic += "    event InputMask(uint[] inpusMaskIndexes);\n";
    textPublic += "    uint public inputmaskCnt;\n";
    textPublic += "    mapping (uint => address) public inputMaskOwner;\n";
    textPublic += "\n";

    textPublic += "    uint public T;\n";
    textPublic += "    mapping (address => bool) public isServer;\n";
    textPublic += "\n";

    textPublic += string(yytext);
    isConstructorHead = true;
}

\) {
    if (isConstructorHead) {
        if (otherParameters) {
            otherParameters = false;
            textPublic += ", ";
        }

        textPublic += "address[] memory servers, uint threshold";
        isConstructorHead = false;
        isConstructorBody = true;
    }
    textPublic += string(yytext);
}

\{ {
    braces++;
    textPublic += string(yytext);
}

\} {
    braces--;
    if (braces == 1 && isConstructorBody) {
        textPublic += "\n";
        textPublic += "        for (uint i = 0; i < servers.length; i++) {\n";
        textPublic += "            isServer[servers[i]] = true;\n";
        textPublic += "        }\n";
        textPublic += "        T = threshold;\n";
        textPublic += "    ";
        isConstructorBody = false;

        textPublic += string(yytext);

        textPublic += "\n";
        textPublic += "\n";
        textPublic += "    function reserveInput(uint num) public {\n";
        textPublic += "        address user = msg.sender;\n";
        textPublic += "        uint[] memory inputMaskIndexes = new uint[](num);\n";
        textPublic += "        for (uint i = 0; i < num; i++) {\n";
        textPublic += "            uint inputMaskIdx = inputmaskCnt++;\n";
        textPublic += "            inputMaskOwner[inputMaskIdx] = user;\n";
        textPublic += "           inputMaskIndexes[i] = inputMaskIdx;\n";
        textPublic += "        }\n";
        textPublic += "        emit InputMask(inputMaskIndexes);\n";
        textPublic += "    }";
    } else textPublic += string(yytext);
}

. {
    if (isConstructorHead) {
        otherParameters = true;
    }
    textPublic += string(yytext);
}

\n {
    textPublic += string(yytext);
}

%%

int main(int argc, char **argv) {

	yylex();

    FILE *solFp;
    solFp = fopen("contracts/tmp.sol", "w");
    fprintf(solFp, "%s", textPublic.c_str());
    fclose(solFp);

	return 0;
}