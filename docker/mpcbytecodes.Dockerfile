FROM sbellem/mpspdz:compiler

COPY src/mpc Programs/Source
COPY scripts/compile.sh .

CMD ["bash", "compile.sh"]
