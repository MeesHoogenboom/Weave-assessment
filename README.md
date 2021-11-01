# Weave-assessment

<!-- Run commands: 
git remote add origin https://github.com/meeshoogenboom/Weave-assessment
go mod init Weave-assessment -->

# Assumptions
1. Data is adhering to the set format (no missing values/data, one row of column titles, comma separated)
2. Data is valid if usage is >= 100 over 2 units of measurement (when a reading is skipped). This can however, be easily changed in the code

# Running the program (WINDOWS)
1. Place the assessment.exe file in the same folder as the .csv file with the data
2. Make sure the .csv file is called data.csv
3. Open a command prompt and navigate to the folder where the .exe and .csv are located
4. Execute assessment.exe [.\assessment.exe]

# Running the program (OTHER OS)
1. Place the assessment.go file in the same folder as the .csv file with the data
2. Open a command prompt and navigate to the folder where the assessment.go and .csv file are located
3. Enter the command [go build assessment.go]
4. Make sure the .csv file is called data.csv
5. Execute the newly compiled assessment program

# Architecture
The program is built upon a simple recursive model that loops over every single line of data and stores it until it reached EOF. It skips the first line (the column headers). At the beginning of each loop it checks whether 1 row of data is already stored in variables so it can compare 2 rows and calculate the usage costs. At the end of each loop, the data of the 'current' row gets stored in variables for use in the next loop. Should the usage over two rows be invalid, the last valid row is stored instead of the new one. When the program sees a new Metering ID the total usage costs gets written to an output.csv file and the cost gets reset. This process is repeated until the writer returns and EOF error. 

# Improvements
1. Better apply Go Conventions
2. Make use of retriever functions, pointers, and structs to increase performance and readability by decreasing the amount of variables needed
3. Adopt test-driven development to make writing unit-tests easier and guarantee code quality/data validity
4. Make use of 'local' packages to reduce clutter in the main Go file

# How I wrote this program
I initially started this project by identifying some key components needed for the basic program to function: 1. A way of opening and reading a .csv file, 2. Logic to compare and store two rows of data, 3. A way to check of the readings were valid, 4. A way to calculate costs based on usage, and 5. A way to write the results to file.
Where I started with a way to read .csv files it quickly expanded by me experimenting and adding on to the original function with the basic knowledge I had of GO. This way of adding functionality without drawing up designs for the program first let me to make some (in hindsight) suboptimal decisions (namely the extensive use of the cvsWriter() function in favour of splitting it up into smaller functions), which made writing tests difficult for me. Unfortunately due to time constraints I was not able to implement extensive tests or restructures, instead I opted to use my time to deliver a functional manually-tested program.