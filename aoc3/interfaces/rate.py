from typing import Annotated

from beartype.vale import Is

PercentageValue = Annotated[int, Is[lambda percent: percent > 0]]
PercentageRate = Annotated[float, Is[lambda rate: rate > 0]]
