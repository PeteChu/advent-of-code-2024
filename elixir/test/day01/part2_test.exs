defmodule Day01.Part2Test do
  use ExUnit.Case, async: true

  test "small input" do
    {:ok, input} = File.read("priv/inputs/day01/input_small.txt")
    expected_output = 31
    assert Day01.Part2.solve(input) == expected_output
  end

  test "large input" do
    {:ok, input} = File.read("priv/inputs/day01/input.txt")
    expected_output = 22_776_016
    assert Day01.Part2.solve(input) == expected_output
  end
end
