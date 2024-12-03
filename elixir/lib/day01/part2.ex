defmodule Day01.Part2 do
  alias Day01

  def solve(input) do
    input
    |> String.split("\n", trim: true)
    |> Stream.map(&Day01.Part1.parse_numbers/1)
    |> Stream.reject(&Enum.empty?/1)
    |> Enum.reduce({[], %{}}, &accumulate_ends/2)
    |> sum
  end

  def accumulate_ends(numbers, {lefts, map}) do
    {[hd(numbers) | lefts], map |> Map.update(List.last(numbers), 1, &(&1 + 1))}
  end

  def sum({lefts, map}) do
    Enum.reduce(lefts, 0, fn x, acc ->
      acc + x * Map.get(map, x, 0)
    end)
  end
end
