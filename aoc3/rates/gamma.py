from dataclasses import dataclass

from rates.rate import Rate


@dataclass
class GammaRate(Rate):
    binary: str
    type: str = "gamma"

    def __init__(self, binary: str):
        self.binary = binary
        self.percentage_value = int(self.binary, 2)
