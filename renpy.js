//contains the graphviz file as string
var graph = "";

function printMessagePromise(msg, boolChoice1, boolChoice2) {
  return new Promise((resolve, reject) => {
    printMessage(msg, boolChoice1, boolChoice2, (err, message) => {
      console.log("cb", msg, err);
      if (err) {
        reject(err);
        return;
      }
      resolve(message);
    });
  });
}

async function getRenpy(repoName, subPath) {
  var errorMessage = document.getElementById("errorMessage");
  errorMessage.style.visibility = "hidden";

  var mainResponse;
  var renpyString = "";
  console.log("fetching start");
  if (subPath) {
    mainResponse = await fetch(
      "https://api.github.com/search/code?accept=application/vnd.github.v3+json&q=label+path:" +
        subPath +
        "+extension:rpy+repo:" +
        repoName
    );
  } else {
    mainResponse = await fetch(
      "https://api.github.com/search/code?accept=application/vnd.github.v3+json&q=label+extension:rpy+repo:" +
        repoName
    );
  }

  if (mainResponse.status != 200) {
    errorMessage.style.visibility = "visible";
  }
  mainAns = await mainResponse.json();

  console.log("fetching end", mainResponse, mainAns);

  for (const item of mainAns.items) {
    // console.log(item.path);
    if (
      !item.path.includes("tl/") &&
      !item.path.includes("options.rpy") &&
      !item.path.includes("gui.rpy") &&
      !item.path.includes("screens.rpy") &&
      !item.path.includes("00")
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

function getRepoStruct(s) {
  const regex = /\w*\/\w*\//g;

  if (regex.test(s)) {
    last = regex.lastIndex;
    console.log(s.substring(0, last), s.substring(last));

    return [s.substring(0, last), s.substring(last)];
  } else {
    console.log(s, null);
    return [s, null];
  }
}

async function getRepo() {
  const loader = document.getElementById("loader");
  loader.style.visibility = "visible";

  var [repoName, subPath] = getRepoStruct(
    document.getElementById("repo").value.trim()
  );
  console.log(repoName, subPath);
  renpyTextList = await getRenpy(repoName, subPath);
  graph = await printMessagePromise(
    renpyTextList,
    document.getElementById("choices").checked,
    !document.getElementById("hideAtoms").checked
  );

  console.log(graph);

  d3.select("#graph").graphviz().renderDot(graph);
  loader.style.visibility = "hidden";
}
