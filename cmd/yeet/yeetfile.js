go.install();

["amd64", "arm64"].forEach(goarch => rpm.build({
    name: "yeet",
    description: "Yeet out actions with maximum haste!",
    homepage: "https://within.website",
    license: "CC0",
    goarch,

    build: (out) => {
        go.build("-o", `${out}/usr/bin/`);
    },
}));