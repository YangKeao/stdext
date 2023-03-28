TEXT ·readTLS(SB),$0-8
    RDFSBASEQ BX
    MOVQ BX, ret(FP)
    RET
