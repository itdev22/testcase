function main() {
  let total = 0;
  for (let i = 0; i <= 6; i++) {
    for (let j = 0; j <= i; j++) {
      if ((i + j) % 2 == 0) {
        total += j * 2;
      } else {
        total += j;
      }
    }
  }
  console.log(total);
}

function deletearaay() {
  const firstArray = [132, 203, 304, 400, 201];
  let newArray = [];
  for (const element of firstArray) {
    if (element == 400) {
      continue;
    }
    newArray.push(element);
  }
  console.log(newArray);
}

deletearaay();
