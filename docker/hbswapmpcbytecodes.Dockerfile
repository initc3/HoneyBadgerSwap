FROM sbellem/mpspdz:compiler

COPY src/mpc Programs/Source
COPY scripts/compile.sh .
COPY scripts/beep.py .

RUN bash compile.sh

CMD ["python", "beep.py"]
