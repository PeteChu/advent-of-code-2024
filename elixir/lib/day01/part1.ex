defmodule Day01.Part1 do
  def solve(input) do
    input
    |> String.split("\n", trim: true)
    |> Stream.map(&parse_numbers/1)
    |> Stream.reject(&Enum.empty?/1)
    |> Enum.reduce({[], []}, &accumulate_ends/2)
    |> calculate_distance
    |> Enum.sum()
  end

  defp parse_numbers(line) do
    line
    |> String.split(" ", trim: true)
    |> Enum.map(&String.to_integer/1)
  end

  defp accumulate_ends(numbers, {lefts, rights}) do
    {[hd(numbers) | lefts], [List.last(numbers) | rights]}
  end

  defp calculate_distance({lefts, rights}) do
    Enum.zip(Enum.sort(lefts), Enum.sort(rights))
    |> Enum.map(fn {l, r} -> abs(l - r) end)
  end
end
