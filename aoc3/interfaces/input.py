from typing import Annotated

from beartype.vale import Is

CleanedInput = Annotated[list[str], Is[lambda lst: bool(lst)]]
