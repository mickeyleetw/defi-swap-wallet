const hre = require("hardhat");

async function main() {
  const initialSupply = hre.ethers.parseUnits("1000", 18); // 1000 MTK
  const Token = await hre.ethers.deployContract("MyToken", [initialSupply]);

  console.log("Token deployed to:", await Token.getAddress());
}

main().catch((error) => {
  console.error(error);
  process.exit(1);
});
