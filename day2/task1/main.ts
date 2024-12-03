function isASafeReport(report: number[]): boolean {
  let isAscending: boolean;

  if (report[0] > report[1]) {
    isAscending = true;
  } else {
    isAscending = false;
  }

  // Start from 1 so we don't have to check for undefined report[i - 1]
  for (let i = 1; i < report.length; i++) {
    if (
      isAscending &&
      report[i - 1] > report[i] &&
      report[i - 1] - report[i] <= 3
    ) {
      continue;
    } else if (
      !isAscending &&
      report[i - 1] < report[i] &&
      report[i] - report[i - 1] <= 3
    ) {
      continue;
    } else {
      return false;
    }
  }

  return true;
}

function main() {
  const contents = Deno.readTextFileSync("../input.txt");

  const lines = contents.split("\n").filter((line) => line !== ""); // filter out the last element that is appended at EOF

  const reports = lines.map((line) => {
    const newArr: number[] = [];
    const numbers = line.split(" ");

    for (const number of numbers) {
      newArr.push(parseInt(number, 10));
    }

    return newArr;
  });

  let safeReports: number = 0;

  for (const report of reports) {
    if (isASafeReport(report)) safeReports++;
  }

  console.log(safeReports);
}

main();
