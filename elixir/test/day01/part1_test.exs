defmodule Day01.Part1Test do
  use ExUnit.Case, async: true
  alias Day01

  test "small input" do
    {:ok, input} = File.read("priv/inputs/day01/input_small.txt")

    expected_output = 11
    assert Day01.Part1.solve(input) == expected_output
  end

  test "large input" do
    {:ok, input} = File.read("priv/inputs/day01/input.txt")

    expected_ouput = 1_660_292
    assert Day01.Part1.solve(input) == expected_ouput
  end
end
