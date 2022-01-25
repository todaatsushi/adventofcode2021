from __future__ import annotations

from dataclasses import dataclass
from typing import TYPE_CHECKING

from beartype import beartype
from interfaces import CleanedInput

if TYPE_CHECKING:
    from typing import Optional


@dataclass
class Input:
    raw_input: str
    clean_input: Optional[CleanedInput] = None

    @beartype
    def get_clean_input(self) -> CleanedInput:
        inputs = self.raw_input.split("\n")
        map(lambda inpt: inpt.strip(), inputs)

        self.clean_input = [inpt for inpt in inputs if inpt]
        return self.clean_input
