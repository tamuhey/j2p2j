package main

const VersionHeader = "# j2p2jVersion: " + Version + "\n"

const CodecellHeader = "# In []:\n"
const MarkdowncellHeader = "# Markdown:\n"
const AuxHeader = "# Aux:"

var Headers = []string{CodecellHeader, MarkdowncellHeader, AuxHeader}

const CellMetaPrefix = "# meta:"

const CelltypeCode = "code"
const CelltypeMarkdown = "markdown"

const modej2p = "j2p"
const modep2j = "p2j"

var modes = []string{modej2p, modep2j}
