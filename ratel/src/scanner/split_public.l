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
    textPublic += "    uint public inputMaskCnt;\n";
    textPublic += "    mapping (uint => address) public inputMaskOwner;\n";
    textPublic += "\n";

    textPublic += "    event NewServer(address server);\n";
    textPublic += "    uint public N;\n";
    textPublic += "    uint public T;\n";
    textPublic += "    mapping (address => bool) public isServer;\n";
    textPublic += "    mapping (address => uint) public votes;\n";
    textPublic += "    mapping (address => mapping (address => bool)) voted;\n";
    textPublic += "\n";

    textPublic += "    uint public opCnt;\n";
    textPublic += "    mapping (uint => string) public opEvent;\n";
    textPublic += "    mapping (uint => bytes) public opContent;\n";
    textPublic += "\n";

    textPublic += "    event GenInputMask(uint inputMaskCnt, uint committeeChangeCnt);\n";
    textPublic += "    uint public isInputMaskReady;\n";
    textPublic += "    uint public committeeChangeCnt;\n";
    textPublic += "    mapping (address => uint) public numCommittee;\n";
    textPublic += "\n";

    textPublic += "    mapping (uint => uint) public voteFinalization;\n";
    textPublic += "    mapping (uint => mapping (address => bool)) public votedFinalization;\n";
    textPublic += "    mapping (uint => uint) public finalizedTime;\n";
    textPublic += "\n";
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

        textPublic += "        N = servers.length;\n";
        textPublic += "        isInputMaskReady = N;\n";
        textPublic += "        require(T <= N / 2);\n";
        textPublic += "        T = threshold;\n";
        textPublic += "        for (uint i = 0; i < servers.length; i++) {\n";
        textPublic += "            isServer[servers[i]] = true;\n";
        textPublic += "            votes[servers[i]] = N;\n";
        textPublic += "            for (uint j = 0; j < servers.length; j++) {\n";
        textPublic += "                voted[servers[i]][servers[j]] = true;\n";
        textPublic += "            }\n";
        textPublic += "        }\n";
        textPublic += "    ";
        isConstructorBody = false;
        textPublic += string(yytext);
        textPublic += "\n";
        textPublic += "\n";
        textPublic += "\n";

        textPublic += "    function registerServer() public {\n";
        textPublic += "        address server = msg.sender;\n";
        textPublic += "        require(!isServer[server]);\n";
        textPublic += "        emit NewServer(server);\n";
        textPublic += "    }\n";
        textPublic += "\n";
        textPublic += "\n";

        textPublic += "    function addServer(address s) public {\n";
        textPublic += "        address server = msg.sender;\n";
        textPublic += "        require(isServer[server]);\n";
        textPublic += "        require(!voted[server][s]);\n";
        textPublic += "        voted[server][s] = true;\n";
        textPublic += "        votes[s] += 1;\n";
        textPublic += "        if (votes[s] > T && !isServer[s]) {\n";
        textPublic += "            isServer[s] = true;\n";
        textPublic += "            N += 1;\n";
        textPublic += "            isInputMaskReady = 0;\n";
        textPublic += "            committeeChangeCnt += 1;\n";
        textPublic += "            emit GenInputMask(inputMaskCnt, committeeChangeCnt);\n";
        textPublic += "        }\n";
        textPublic += "    }\n";
        textPublic += "\n";
        textPublic += "\n";

        textPublic += "    function setReady(uint num) public {\n";
        textPublic += "        address server = msg.sender;\n";
        textPublic += "        require(isServer[server]);\n";
        textPublic += "        if (numCommittee[server] < num) {\n";
        textPublic += "            numCommittee[server] = num;\n";
        textPublic += "            isInputMaskReady += 1;\n";
        textPublic += "        }\n";
        textPublic += "    }\n";
        textPublic += "\n";
        textPublic += "\n";

        textPublic += "    function resetThreshold() public {\n";
        textPublic += "    }\n";
        textPublic += "\n";
        textPublic += "\n";

        textPublic += "    function removeServer(address s) public {\n";
        textPublic += "        address server = msg.sender;\n";
        textPublic += "        require(isServer[server]);\n";
        textPublic += "        require(voted[server][s]);\n";
        textPublic += "        voted[server][s] = false;\n";
        textPublic += "        votes[s] -= 1;\n";
        textPublic += "        if (votes[s] <= T) {\n";
        textPublic += "            isServer[s] = false;\n";
        textPublic += "            N -= 1;\n";
        textPublic += "        }\n";
        textPublic += "    }\n";
        textPublic += "\n";
        textPublic += "\n";

        textPublic += "    function reserveInput(uint num) public {\n";
        textPublic += "        require(isInputMaskReady > T);\n";
        textPublic += "        address user = msg.sender;\n";
        textPublic += "        uint[] memory inputMaskIndexes = new uint[](num);\n";
        textPublic += "        for (uint i = 0; i < num; i++) {\n";
        textPublic += "            uint inputMaskIdx = inputMaskCnt++;\n";
        textPublic += "            inputMaskOwner[inputMaskIdx] = user;\n";
        textPublic += "           inputMaskIndexes[i] = inputMaskIdx;\n";
        textPublic += "        }\n";
        textPublic += "        emit InputMask(inputMaskIndexes);\n";
        textPublic += "    }\n";
        textPublic += "\n";
        textPublic += "\n";

        textPublic += "    function finalize(uint seq) public {\n";
        textPublic += "        address server = msg.sender;\n";
        textPublic += "        require(isServer[server] && !votedFinalization[seq][server]);\n";
        textPublic += "        votedFinalization[seq][server] = true;\n";
        textPublic += "        voteFinalization[seq] += 1;\n";
        textPublic += "        if (voteFinalization[seq] > T && finalizedTime[seq] == 0) {\n";
        textPublic += "            finalizedTime[seq] = block.number;\n";
        textPublic += "        }\n";
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