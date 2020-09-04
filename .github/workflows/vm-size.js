var fs = require("fs");
console.log("Going to write into existing file");

var goCodeFilePath = "azure_vm_size_list.go";

// Write headers and static code at top of azure_vm_size.go file
fs.writeFile(goCodeFilePath, 'package rules\n\nvar azureVmSizeList = []string{\n', function(err) {
   if (err) {
      return console.error(err);
   }
   console.log("File header written successfully.");
});

// Get data from vm-size.json file
const jsonFileData = fs.readFileSync('vm-size.json', 'utf-8');
var dataArray = JSON.parse(jsonFileData);

// Write vm size data to azure_vm_size.go file
for (i in dataArray) {
    fs.appendFile(goCodeFilePath, "    \"" + dataArray[i] + "\",\n", function(err) {
        if (err) {
            return console.error(err);
        }
    });
    console.log("VM sizes populated in file successfully.");
}

// Write end of azure_vm_size.go file
fs.appendFile(goCodeFilePath, '}', function (err) {
if (err) {
    return console.error(err);
}
});
console.log("End of file written successfully.");

