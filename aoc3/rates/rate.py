from dataclasses import dataclass

from beartype import beartype
from interfaces import PercentageRate, PercentageValue


@dataclass
class Rate:
    percentage_value: PercentageValue

    @beartype
    def as_rate(self) -> PercentageRate:
        return self.percentage_value / 100

