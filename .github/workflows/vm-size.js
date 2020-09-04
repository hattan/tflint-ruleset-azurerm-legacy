var fs = require("fs");
console.log("Going to write into existing file");

var goCodeFilePath = "azure_vm_size_list.go";

// Write headers and static code at top of go code file
fs.writeFile(goCodeFilePath, 'package rules\n\nvar azureVmSizeList = []string{\n', function(err) {
    if (err) {
        return console.error(err);
    }
    console.log("File header written successfully.");

    // Get data from vm-size.json file
    const jsonFileData = fs.readFileSync('vm-size.json', 'utf-8');
    var dataArray = JSON.parse(jsonFileData);

    // Create list of vm size elements to write out to file
    var vmSizeString = "";
    for (i in dataArray) {
        vmSizeString += ("    \"" + dataArray[i] + "\",\n");
    }
    console.log("vmSizeString:");
    console.log(vmSizeString);

    // Write vm size data to go code file
    fs.appendFile(goCodeFilePath, vmSizeString, function(err) {
        if (err) {
            return console.error(err);
        }
        console.log("VM size data written successfully.");

        // Write end of go code file
        fs.appendFile(goCodeFilePath, '}', function (err) {
            if (err) {
                return console.error(err);
            }
            console.log("End of file written successfully.");
        });
    })
});
