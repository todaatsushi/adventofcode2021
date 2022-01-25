import unittest

from adapter import InputAdapter


class InputTestCase(unittest.TestCase):
    def setUp(self):
        self.TEST_INPUT = (
            "00100\n"
            "11110\n"
            "10110\n"
            "10111\n"
            "10101\n"
            "01111\n"
            "00111\n"
            "11100\n"
            "10000\n"
            "11001\n"
            "00010\n"
            "01010\n"
        )
        self.TEST_OUTPUT = (
            "00100",
            "11110",
            "10110",
            "10111",
            "10101",
            "01111",
            "00111",
            "11100",
            "10000",
            "11001",
            "00010",
            "01010",
        )

    def test_input_adapter(self):
        adapter = adapters.Input(
            raw_input=TEST_INPUT
        )

        cleaned_output = adapter.get_clean_input()
        self.assertEqual(len(cleaned_output), len(self.TEST_OUTPUT))
        for index, output in enumerate(cleaned_output):
            with self.subTest(output=output):
                self.assertEqual(
                    output,
                    self.TEST_OUTPUT[index]
                )

if __name__ == "__main__":
    unittest.main()
