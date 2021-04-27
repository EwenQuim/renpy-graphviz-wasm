//contains the graphviz file as string
var graph = "";

function printMessagePromise(msg) {
  return new Promise((resolve, reject) => {
    printMessage(msg, (err, message) => {
      console.log("cb", msg, err);
      if (err) {
        reject(err);
        return;
      }
      resolve(message);
    });
  });
}

async function getRenpy(repoName) {
  var errorMessage = document.getElementById("errorMessage");
  errorMessage.style.visibility = "hidden";

  renpyString = "";
  console.log("fetching start");
  const mainResponse = await fetch(
    "https://api.github.com/search/code?accept=application/vnd.github.v3+json&q=extension:rpy+repo:" +
      repoName
  );

  if (mainResponse.status != 200) {
    errorMessage.style.visibility = "visible";
  }
  mainAns = await mainResponse.json();

  console.log("fetching end", mainResponse, mainAns);

  for (const item of mainAns.items) {
    // console.log(item.path);
    if (
      !item.path.includes("tl") &&
      !item.path.includes("options.rpy") &&
      !item.path.includes("gui.rpy") &&
      !item.path.includes("screens.rpy")
    ) {
      rawFileUrl = item.html_url
        .replace("github.com", "raw.githubusercontent.com")
        .replace("blob/", "");
      // console.log(rawFileUrl);
      const rep = await fetch(rawFileUrl);
      const ans = await rep.text();
      renpyString = renpyString
        .concat(ans)
        .concat("\n#renpy-graphviz: BREAK\n");
    }
  }

  return renpyString;
}

async function getRepo() {
  const loader = document.getElementById("loader");
  loader.style.visibility = "visible";

  var repoName = document.getElementById("repo").value.trim();
  console.log(repoName);
  renpyTextList = await getRenpy(repoName);
  graph = await printMessagePromise(renpyTextList);

  console.log(graph);

  d3.select("#graph").graphviz().renderDot(graph);
  loader.style.visibility = "hidden";
}
