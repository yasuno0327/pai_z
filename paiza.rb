strings = gets.split.map(&:to_i)
limit_zero = strings[1]

input_string = gets
question = input_string.count("?")
if input_string.match(/[?0-1]/) && strings[0] <= 10
  for i in 0...2**question do
    string = input_string.clone
      i.to_s(2).chars do |char|
        for j in 0...string.length do
          if string[j] == '?'
            string[j] = char
            break;
          end
        end
      end
    string.tr!("?", "0")
    if string.count("0") < limit_zero
      puts string
    end
  end
else
  puts 'None'
end
