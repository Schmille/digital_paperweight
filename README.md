# Digital paperweight creator
> Congratulations Matt, you win a pair of digital paperweights!
> 2 Gigabytes each, they do the exact same thing as regular paperweights, which is they take up space utterly uselessly!
> Ahem... so... the're on this memory stick, you can copy them to your desktop... ahm... when you need them!
>  
>‐ Tom Scott, 2012 

Inspired by and based on the excellent price from [this](https://www.techdif.co.uk/podcast/05-weshouldprobablyhaveawomanonthisshow.mp3) episode of the Technical Difficulties'
reverse trivia podcast.

## Compiling
You can compile this project, as you would any other Go project.

```
git clone https://github.com/Schmille/digital_paperweight.git
cd digital_paperweight
go build .
```

## Usage
After compiling the project, you can simply run the executable from terminal. 
\
The desired paperweight size must be entered as the first (and only) commandline argument.

For example `.\digital_paperweight.exe 2GB` on Windows.

Acceptable size units are: 
 - GB
 - MB
 - KB
 - B

**NOTE**

If you do not have enough RAM to contain the paperweight (plus a safety margin of 2GB), the generator will use a 
stream writing mode which will write the paperweight to disk in 128-bit chunks. This will save memory, but will 
also cause significantly more write operations, thus greatly increasing the required time.

It is generally recommended, that you use reasonable sizes.