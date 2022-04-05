PaperMC:
let versionNum;
let buildNum;

const Name = "";
const Description = "";
const Homepage = "";
const License = "";
const sha256 = "";
const URL = liveCheck();

function liveCheck() {
    return `https://papermc.io/api/v2/projects/paper/versions/${versionNum}/builds/${buildNum}/downloads/paper-${versionNum}-${buildNum}.jar`
}

