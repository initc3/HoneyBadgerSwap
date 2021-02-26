FROM sbellem/mpspdz:compiler

COPY src/mpc Programs/Source
COPY scripts/compile.sh .

RUN bash compile.sh
