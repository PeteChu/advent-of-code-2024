defmodule Day01.Part1 do
  def solve(input) do
    x =
      input
      |> String.split("\n")
      |> Enum.map(fn x ->
        x
        |> String.split(" ", trim: true)
        |> Enum.map(fn x -> String.to_integer(x) end)
      end)
      |> Enum.filter(fn x -> length(x) > 0 end)

    l = left_list(x) |> Enum.sort()
    r = right_list(x) |> Enum.sort()

    Enum.zip(l, r)
    |> Enum.reduce(0, fn {l, r}, acc -> acc + abs(l - r) end)
  end

  defp left_list(input) do
    input |> Enum.map(fn x -> List.first(x) end)
  end

  defp right_list(input) do
    input |> Enum.map(fn x -> List.last(x) end)
  end
end
