package main

import (
	"time"
)

func main() {
	start := time.Now()
	toMatrix := [][]string{}
	line := []string{}

	data := realInput()
	lenData := len(data)
	for i, letter := range data {
		if letter == '\n' {
			toMatrix = append(toMatrix, line)
			line = []string{}
			continue
		}

		line = append(line, string(letter))

		if i == lenData-1 {
			toMatrix = append(toMatrix, line)
			break
		}
	}

	count := 0
	for i := 0; i < len(toMatrix); i++ {
		for j := 0; j < len(toMatrix[i]); j++ {
			// Horizontal
			if j < len(toMatrix[i])-3 {
				if checkLetters(toMatrix[i][j], toMatrix[i][j+1], toMatrix[i][j+2], toMatrix[i][j+3]) {
					count++
				}
			}

			// Vertical
			if i < len(toMatrix)-3 {
				if checkLetters(toMatrix[i][j], toMatrix[i+1][j], toMatrix[i+2][j], toMatrix[i+3][j]) {
					count++
				}
			}

			// Diagonal right
			if i < len(toMatrix)-3 && j < len(toMatrix[i])-3 {
				if checkLetters(toMatrix[i][j], toMatrix[i+1][j+1], toMatrix[i+2][j+2], toMatrix[i+3][j+3]) {
					count++
				}
			}

			// Diagonal left
			if i < len(toMatrix)-3 && j > 2 {
				if checkLetters(toMatrix[i][j], toMatrix[i+1][j-1], toMatrix[i+2][j-2], toMatrix[i+3][j-3]) {
					count++
				}
			}
		}
	}

	duration := time.Since(start)
	if count == 0 {
		println("XMAS not found")
	} else {
		println("XMAS found", count, "times")
	}
	println("Execution time", duration.Microseconds(), "µs")
}

func checkLetters(letter1, letter2, letter3, letter4 string) bool {
	if letter1 == "X" && letter2 == "M" && letter3 == "A" && letter4 == "S" {
		return true
	}

	if letter1 == "S" && letter2 == "A" && letter3 == "M" && letter4 == "X" {
		return true
	}

	return false
}

func input() string {
	return `....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`
}

func realInput() string {
	return `AAMXMASASMXSAMXXSAMSAMXXSXMMSMMSXSASXSXSSSXSAMXXXSXMASXMAMMASAMXAXAXMXMSSMMSXSASXSAMXXMASMXSMMMMMXXMAXXSXXXAXMSAMXSMXMSXSAMXMXMSXMMXSMXMXSMS
AXMASXMAASAMMMSMSAMSAMXMAXMAAAMXMASAAMMMAAAMMXMASMAMAMASMSAASXMAXMMXASAMXAAAMSAMMMAMMMXMASASAXXASMMMMSAMMMXMASXMASMSMMMAMMMMMAXXAMXAXMAMSAAA
XMSMMAMXMMSMAAAAMAMSAMXAMAMSSSMXAMMMSMAMMMMMSASXAXAMASXMAAMXXMSMXSAMXSMMSMMSMMAMASAMMAAXSMMSAMSMMAAAMMAMAASMMMXMMXMASAMAMSMSSSSSSMMXSXMXAMXM
SAMXSAMSXMASMSSMSMMMMMMAMMMMAMASMXSAAMSMMAMXSASMSSXSASMMSMMMSMAAAXAAAXAAMAMXASMMMSASXSMSMMXMXMAXMSMSSSSMSMXAASMMAASAMXSASAAXAAMXAASMMASAMXAX
ASAASXMMAXAXMAXXAXAXASAMXAAXMMMMMAMXSXMASXSMMMMAAMXMASAAMXSAAMMAMSSMSMMMSSMSMMXMXSXMAXXAAXMMSSMSXMAMAAAAXMSSMXASXXMMSMSMSMSMMMMSSMMASAMAXSSS
MMMMSAXSAMMXSASMMSMSAXMASMSMXAXAMXMXXASAMXAMXXXMAXAMXSMMMXMXSMXAXMAAXXSAAXMAMXXMMSAMXMSSXMAAMAAAMMAMMMMMMMAMAMXMXAXXAASAMXAAXXXXMAMMMMSXMAAX
XMAMXMXMASXAMMSAXAAMAMSASMXXSXSXSAMXMAMSMSAMXSMSASXSAMXSMASAXASMMSMMMAMMSSSMSMSMASXMAMMAMSMMSMMMMSASXXAAAMAXXMXASMMMMSMAMSMSMMMXMMMSXMSMSMAM
SSMSAXSMASMMMASMMMMMXMMASAMASAAMSASAMXMAMMAXXAAAAMASXSASMXMASAXXAAAAMAMAAAMXAASMAMASMSMAMXAAMXSXASASASXSSSSMSXMASXAAXAXSMSMAAAMMASASAAXAXXAX
AAAMXMXMAXAXMXXXXXXXXXMMMXMAMXMXMASXMXSASMSMXMSMSMXMAMASAMMMMXMMSSSMSXSXMMSXMSSXMSMMMAMMSMMMSAMMMMAMAMAAAAXAXMXAXMSMSAMXAMMSSMAAXMASXMMSMSAS
MMMMSSXMASMXMASMMMXSXMMMAXMSMSXAMMMXMASASMMMAAAMXMXMXMXMMMXAMASXMAMAXMMMXMAASAXXXAXASASMSXSAMMSAXMAMAMMMMAMMMAMMMXAXMMMMAMAAAMXSXMXMAMAXXSSM
XMAAMAAMAMMAXAAAAXAMMAAMAMMMAXMXSAAAMAMMMAAXXMMMAMAXMASMMSMASAMMMMMMMAAAMMSSMASMSMSXSAMAMSMASXSMSSSSSMSAMXMXASAMXMSMMSASASMMXXMMMXAXAMMSMMAS
MSMSSSSMMSXMMMSSXMMSASXSSMSMMMMMMMSMMASASMMSSMXSASXXSASAAXSXMMXMAAMXSSSXSAXXMAMXAAAXMAMMMASMMMXAXAMAAASXSMXASXXAAAAAXSXSASAMSAAAAMMSXXXAAXAM
XAAAAAXXAAMSXMAMXXMAMMAAAAAAAAAAXAXAXXXXMAAAAAASAMAAMASMMXAMXSASMMSAAAAAMXMXMXSSMSMXMAMAMAMMAMMMMAMXMMMMMAMXXMSSMSSSMSMMAMAAASXMSXMAASMSSMXS
SMSMMMMSMXMXAMASMMSSMMMMMMMXMSSSMSSMMSSSSMMMMMMSAMMMMAMAXSAMXSASAAMMMMMMMMMXMAAAAAXSSMXSMSSSMMAAMAMXXAAAMXXAMXAAMMAAXXAMXMMMMMXMAASMMMAAXXXX
SAMXMAMSMSMXAMAMAAXASAAXXXSAAAAAAAXMAXAAMXXXXAXMASXAMASXMMASXMAMMMMSASXMAAXAMXXMSMSAAXAMXMAAASXMSMSASMSMSAMSSMSSMMSMMSSMAXXAXXAMSXMAAMMMMMSM
SAMXSASXASAASMSMMMMMSXXSAASMMMSMMMSMSMMMMXMMSMSXAXXXSAMXASAMAXSXXXASMSASXSSSXSXXAXMMMMMSAMMMMSAAXAAXAAMAMMSAAAAXAXAXAXMMSMSAXSMMMMSSMSXSXAAA
MAMXSSSMXMAMMAAXSXMXXXXMXMMASAXXSXAAXMMMXMXXAMXMMSSMMMMSXMASMMMMMMAMASXMXAAAASXXXXXAMAMMAMXSASMMMXMSMMMAMSMXSMMSSMMMSSSXMASXMXMAAMAAAXAXMASM
SSMMMAXAMXMMMMMMMAXAXSSMMXSAMXMMSSMMMAAXAXXSSSSXSAMXSAAMMXAMMXMAMMSMAMXSSMMMMMAXSMSSSSXSASXMASAMSAMXAASMSSXMMAXXMAAAMAMAMXMMMASXSSSMXMXMAMXM
MASXMMMXAAMAXMASMMMXSAAAAXMAMXMAMAXXSXMXMXMXMAMSMASAMMSMAMSMSAMMSAAMAMAXAAXXXMSMAAAAAAASAMXMXMAMSXSSXMASAXAAXMXSSSMSMASXMAXAMXSAMAMXSMMSXMAX
SAMXMXMXSXSAMAAAAAAXAMMMMXAXMAMSSSSMXAMASAMAMAMAXSMXSAXMAXMASXSAMXSMMMSSMSXSMMASMXMMMMMMAMMSASAMXAAMMSMMMSMMXMAMAXMXXMMMXMMSMMMAMXMASXAAAMXM
MASASASMMASMSMMXSMMXMXMMMMMMMXSXAAXMAXSASASASXSMMMAXMAMSSSMAMAMMSAMAXAXXXXAAASASXSXAAXXMAAMSXSASMSMMMMMAMMAMXMASMSMSMSMSASMXAASAMAXMXSXSAMXA
MXSASAMAXAXXAXMXMAMXMAXSAAXAXMMMMMMMXXMASXAASAAXAMXMXAXXAXMAMAMXMAMSMMXSAMMMMMMSASXSXSXMASXMASXMXAMXAAMAXSAMXSASXAAAAAXXAMASMMMASMSXAMMAMSSS
XXMMMMSSMMSSMXXAXMAXXAXXXMXMSAMMMXMSXAMAXXMSMMMMXXMASMXXMASMSXSASXAMAMAMASMSMAXMXMXMXMAXAMAMAMXMASXSSSMSXSXSXSXSXMSMSMMMAMAMAASMAMMMASAMXAAM
XMASAMXMAXAAAAMSMSMSMSXSASAXSAMAXSAMSXMASMXAASXMXSXXAXSXXXMASASASAMSAMXSAMXASMSMXMASXMXSASXMMXMXAMXMAAAMASMSXXAMMXAXXAXSMMMSXMMSXAMMMSXSMMXM
XSAMXSASMMMSMMMXAAAXAAAXASMXXSMAMMAMXXXXXAMMXMASAMXSXMASXXXXSXMXMAXSAMAMASXMMMAMAMAAAAASAMXXSASXSSSMMMMMASAMMMMMMXAMXMMXAAASASMMSXASASASASAS
MMASXMASXAXAXXAMSMSMSMAMXMXXMXMMSSSMSSSMXMAXASMMXXMXAMMMMMSASXSASAMXAMSXMMAAASASXMASMMMMAXAXXASXAAXXXMXMAMMMAXMAMMSMXMAMMMXXAMAAASMMASASAMAA
AXAMAMAMMSXMSMSMAXXAXXMAXMMMXXAMAMAMAAAXSXXSAMAMSMSAMMXSAAMMMAXASAXXSAXAXSSMMSMSASXXMASXSMMSMSMMMMMMAXAMAXMXSXMAMMXAAMMMSMSMASXMMAXMMMAMMMSM
SMSSMMASXMAXXAAXMSMXMAMSMXAASMSMASMMMSMMXAAXASMMAAAAMSXMMMXAMAMXMXSAMASAMXXAAXASAMXMSASMXAAAXAAXMAAXXSXSASMMMASMSSMMMSXAAAMSAMASXMASMMSMSAMX
XMAAASXSASAMMSMXSAXAXXMAXSMSMAASAMAAAAXMMSMSAMASMXMSMXAMASXMMASAAXMXMMMMMMSMMSXMMMAAMASXSSSMSSSMMMSMXXAXAAAASAMXAAXSAMXXMMMMXSAMXXAMAMMAMASA
MMMSMMXMMMMMXXAXSAMMMSSMXXXXMSMMSSSMSSSMAAXMASMMXMSAMSXMAXAXSAXMSMMXSAAAMMAAAXAAXSSSMAMXAMXMAXAASAXAXMSMSSSMXAMMMMMMASMSXSASMMASMMMXAMMAMMMM
SAAXAMAMXAMXXMSMMAMXMAASAMXSAMXSAXMAXXMMSSMSAMXMAXMAMSAMSSMMMMMMAAXASMSMSMSSMSSMMXXAMASXMXAXMMSMMASAMAAAMMXMSMMAXXXSAMXAXSAMAXMMXAASXSSSSSXS
MMXSAMAXSASAMXXASAMXMSSMMAAMAMXMASMMMAMXAAAMMSSXMSSMMSAMAAXAASASMSMMSAXXAMMAMMAASXSXMXMASAMXMAMAMMMXASMMMXSMASMMSAMMMXMMMMMMSXMMMMMSXAAAAAAX
XXASAMAMXXMASASMSASXMXMASMSXMASXMAXASAMMMSMSXXMAMMAAAMAMSSMMMSAXAAAXMAMMMXXAMSXMMAMSMASMMMXSMASAMXAXMXXSAMASASMAXXMASAMXAAXAMMMMASXMMXMMMMSS
SMMSXMMMSMSXMASAXMMXMASMMAAXXAXAXMAMMASMAMMMXASXMSSMMSAMXAASMMMMSSMMMAMSXSSMMSMMMXMAMAMAMXXMXXMAMMSXMAXMAMAMXSXAXXSAMASXSSMMSMAAMMAAXMXAXXAX
SAXXMAMAAASMMAMMSXMASASAMMMMMSSSXMMSSXSMMSAAMXAMXAMAMSMSSXMXAAXXAAASXSXSAMXMAMAXAXSAMSSSMSMAMMSSMXXAMXMXMMXXAMXMAMMMSXMMXAMAAXXSXSSMMASXSMMM
SXMASAMSSMMSMXSMAASASMSAMMXSXMAMAAAXMXSAMSMMSAMAMMSSMXAAXMASMMMMXSMMXMAMXMAMAMXMXAMXSMAXAXAAMMAMXASXMAAASMXSASAMMAAXAAMXSAMSMSXMMMASAMXAAMXS
MAMXMAMAAAASXXMMSMMAMAMAMXMMAMAMXMMXXAMAMXSAMXMXSAAASMMMXSAMXAASMMMXSMMMAMMXSSMMXMASAMAMSMXSSMXSMMSASMSMSAASASXSXSMMXAAAMAMXXSASASAMSSMMMMAX
SAMXSSMSSMMSMXXAMMMMSMSSMMAXXMSSXSMXMASAMXMSMSAMMMSSXXMAMMASMXXXAMXAMAASXMSAAAAXASMSAMXSAAAXAMAXMXXAMXAXMMMMXMMAMAAXSMMXSMMXAMAMAMASXSAAAMXS
SASAAAAAAMXSASMMSAAAAXXMASMXMXXMASAASXMASXMAMAAAXXMAMXMASXXMASMSMMMSMSMSMAMMXMMSMSAXXMXMMMMMAMXSMSMSMSXMXXXMASAMXSAMXAMMSAMMSMSMMMXMAMSMSSMS
SAMXSMMSSMAMXSAASMMSMXXMMMSAMSAMMMSMSASXMMSASMMMMMMMSXSMSAMXAAAAASAXAMMXMAMMXSAMXMAMXSAMXAMSXMAAAXAAXAMXMXAMMASXAMAMXXMAXAMAMXMAMMMMXMAXXXAM
MXMXXXMXMMXSAMMMSAMAMSMMSASXMASMAMAXSAMAAAMXSXAAXAMXMASXSXMAMMMSXMAMXMAXXXSAMXASMSXMXAMMSAXXMMSSSMSMSMXSAMXMXMMMXSAMXSSSSSMXXAXMMAASMXXSMMSM
SMMMXMMAAMMMMXXXXAMAXSAXMASXMAMSXSXXMMMMMMXSXMSMSXXAMXMASASXXMAXAMAMMMMSMAMXMSMMAMMMXMMAMSAAXAAMAMXAXAMXSMAMAXXAAXMMXMAMXAASMSSMSXXSAMXAAMAM
AAAAAXSSMSAAMASXMSMSSMSMMSMXMAMMMXMMAMSXAXMXAMXAAMXMMAMMMAMXASAMXXMSAAAAMAMMMSSMAMAXASMXXMSMMSXSAMMMMSMAMXSSSMMMSMMSXMAMMMMMXAAMMXASMASXMSAS
SSMXMXAXAXXMSAMAAMAMAMXMXXAMMMMSMMASXMASMSSMAMMMMSAMSSSXMAMXMMXMMMMMXMSXMASXAMAXXMMMAXAMXMAMXMASXSASAMMXMAXAMAMAAAAAXSASXXMSMSMMMMXSXMAAASAM
MAAASMMMSMSXMASMMMXMAMXMAMXMASAAAAXAMXMMXAASXMSMAXMAAMAASMSASMSAXSAXAMXMXMSMMSAMMASMSMAMSSSMXMAMXSAMASASMSMAMXXSXXMXMMAMXMAXAAAAXMASMSMMXMMS
MMMMXAMAMXMASAMXMASMXMXMASMMAMSXMXSXSAXMMSXMXAAMASXMMSSMMASASAAMMSASXMMSAAXXXMASXAXAAMAMXAAMAMASMMXMSMAMAXMAMXXMMSSMMSSMMMASMSSMSSXMASXMXXXA
SASMSSMASAXMMSAMXMSXSAMSASAMAXAMSMSASAMXAMAMMMMMXAXMAAXAMXMAMAMXAMMMXAASMSMMMSAMMMMMMSAMMMMMAMAMAMXMAMXMAMSSSMSAAAAXAAAAXMASAAXAXXAMSMAMSMXM
SASXAASMMMXXAXMAAXMAMMAMAMXMSMSMSAMXMAMMXSAMSSMSAMXMMMSXMMAXXAXMSXSAMMMMAXXAAMMXMAXAASXXAMAXMMSXSXAXASMMXMMASASAMSSMMXSMMXAMAMMMMXMMXSMXAAAX
MXMMSMMMAMMMSMMSMSMXMAXMSMSMXAAAMMMMSAMXXMAMXAMXMASXAAXAXAMMXAMXXXMASMSMXMMMXSMASASMMMASXSSSMAMAXXASAMXAMXXAMMMAMAXAMAMMMMMSXSAAAASAMAXSMSMS
XSXAAMMMAAAAAAXMAXMMASMAAAAAMSMMMXSASASMMMSSMMMSMAMSMSSSMSAMSMSMSXSAMXAMMXMAXXXAMXAMXMAMXAAXMAMMMMMAMXAMXMMSMXSAMXSXMAMAAAMAASMXSAXAXSMSAAAX
MAMSXSASXSMSSSMMAMAXAAMXMXMMMAAXMASMSAMAAAAAAAAXXAXXSAAAAXAAAAAAAMMAMSXSMASMSMSMMSSSSMMMXMMMMXSAMMXAXMAMAMMMAAAXSMXMMMSSMSMMMMAXMXSMMMAMSMMM
AMXXAMXSXAXAMMMSSMMMMSMXMXMASMSSMASAMMSSMMSSMMXSSMSXMMSMMXMSMSMMMMSMMSXMXMMAAAAAMAAAXAMMXXSXMAMASMMMMXAMMSASMMMXSAMXAAXXAAXXMMMMMASAAMAMAMXA
SXMMMSMSMSMXSAAAAAXXMAXAAASASAAAMAMAXAAAXMAXAXMMAAXMMAXMAMXMAAMXSAMAAXAMASMXMXSSMMSMSXMASMMMMASMMXAAASMSAMXSMSXMMMMSMSSMSMSMMAAAMAMMMMASAMMS
AMXAAAXXAXAAAMMSXMMXSASMSMSAMMSSMMSMMMXXMMASMMMSAMSAMMMMSAMMXMSASAMMMSASASMMSAMXAXXMAXXXAAAMMMSXASMSXMSAMXXMXSAMXXXXAAMAAMXASMXSMMSAXSASASAX
MASMSSSMSMSMXMAXXSAXMASXAASXXXAXXXAMAMXSXMAMXAXAAMXMMXASASMAMSMMMAMXMXMMASAASAMXSMXAASMMSMXMAAMMMMXXAMXMASMSASAMSSMMMMMXMSSMMSMAAMAMXMASXMXS
XMXMAMAMMAMXXXAMXMASMMSAMXMMXMASXSXSMSXAAMSSSSSSSSSXMSMSAXMMSAAXSMMMMASMSMMMSAMAMMXSAMXAAAXSMMSAXSMSXXAMXMASASMMMMAXMASAMAAXAAMSMMXXAMAMAXAX
SAAMMSAMMXMMSAMXSMASAMXXXXXASMMSXMASASMSMMAAAXMXAAMXMAAMXMAMSXSMSAAASASXAAMAXAMAXAMXAMMMMXMAMXMASXAXMSMXXMXMXMXAASMMSASASXMMMSXMSAMXMMSSSMSA
XSXMAXASMAXAXAMXMMMSAMASMMASAAASAMMMAMAAMXMXMAMMMMMMSMSAASXMSXMASMMMMAXMSSMXSXMSSSXSAMXMASMSMSSSSMMMXAAMSMMSSMSXMXAMMMXXMAXAAMAMXMASXMAAAAAM
MAXMSSMMXMSSMMMXXAXMMMASASMXSMMSAMAMAMSMSAXSXMXSAMXAXMXMMXASXMMAMSSSMSMAXXMMSXAXAMASXMMMMSAAAAXAXAAASMSMAAASAMXMSSSMXASXMSMSXMAMMXXSAMMSMMMX
AMXXAAXMAMAAAXAXSMSAAMXMAMXAMXMSAMXMSMMASMSSMMMSAMXSMSMMMSMMAXMSMAXAAAMSMSAAXMMMAMAMAAAXAMXMMMSMSSMMSAAMMMMMMMSXAAAMMMSAAXAMXXXMXMASAMAXAXXM
SXMMSSMSMMSXMMSMMXSXMXMMAMMSSSXXAXAXXASAXMAMAAAXXAXMXSAAAAASMMXXAMSMMMXAASMSSXSSMMASMMXMMSMXXXAXAAAAMXMSXSSSMASMMSMMAXMMMMAMXMSSSMXSAMXSMMMA
XAMXMAMAXAXAMAMAMAMXXAAMAMXMAMAXAMSXSAMXSMMSMMMSASMSASXMSSMMMXSMSMAXMXSMMMAAAXMASMXSAMXSXMMMSSMSSSMMMXXAAAAMMASXMAXSMXMXMSAMASAAXAASMMAMASAS
SAXSAMXXMMMMMAMMMXSAMXSMAMMMAMXSAMXMMXMMXAASMMMMAMAMAXXXAXMASMXAAMXMXAAAXXXMMMSAMXAXAXASASXAMAMAMXMMSSSMMMMMMMSMXMXSMSMAMSXSXMMSMMMSAMXSAXSM
SAMXXSMMXAASMMSXSAMXXAAXXMAXXSAAXSAMXMASXMMSAXXMAMXMSMSMMXSAMXMSMSSMMMSAMXMMXMMAMMXMAMASAMMSSMMMSMXXAAXXAAAMXMSMMSAMXAXAXSMMAAXMAMAXMAXMXMAM
XXAXAAAAMXXSAXSAMASAMSSMMSXSAAMMMSXXSAXXSAAMXMMXXSXAAAAAAAMASXXMMAAAAAMAMMAXASMSMSMXMAXMMMAXXXXXXXMMMSMXSSMSAMXAAMASXMSSMMASMMSSSMASXSMASMMM
MASMSMMMSMAMXMMASXMAMAAAXAMXMMSXAXXMASAAXMMMXXSAAAMSMMMMMXMAMXMAMMSAMXXAMSSSXSAXASAAMSMSASMMSMMMXMASAAAXMXMAMXMMMSXMMAAMSSXMMAXAMXAMXAMAMAMX
AMXAXXMAXMAMAXXAMXMXASXMMMSSXXXMMSMXAMMXMMMSMAAMSMMXMMSMMMMSSMSAMMMMMSMAMXAAMMAMAMXMXAAXASAASAAAASMMSSXMMAXAMXMSAMXMMMMSAXAXSXMXMASAMMMASAMS
XXMXMAMXXSMSMSMMSMAAXASXSAAXSASMMAXMASXASAAAMXMAMXAMXAAAAAMAAXSASXSMAAXAMMSMAMAMMMASMMXMAMMSMSMSXSAMXXAASMMMSAAMASASMMMMMSMMMAMMMMMAMXSASMSA
SSMSMSMMMMAMAAXSMMSMSASMXMASMXMASAMSAMXAMMSMMXXASAMMMXMMMSMMSMSXMMMMSSSMSAMXSSXSMMMSAAMSAMXAAAMMASMMAXSMMAAASMSMSSXSAASAXAAAXAMAAASAMXMASXMM
XMAAMAMAAMAMSMSAXAAAMXMAMXMXMXSAMXMMSSMMSXMASMMXSASXXSXSAAASMMXXAAAAAAMXMMSMAMMMXAAMMXMSASMMSMXSAMXMAMMXMSXMXSAXAXAXMMSMSSSSSSSSMMSXMAAASMSX
MMSMSASMSXXXAMXMMMSXMAMSMASAMMSASXAAAXAAAASAMASMMAMAAAAXSSXMAMMMSSMSMMMXMASXXMASMMSSXMXSAMAMXXAMASXMXMMAXXMASMXXMMASMMMMAAMAAXMASMXMXXSAXAAX
AAMAMXSXMXSSXXAXMXAASMXMMXXASASAMMMMMSMMMMMAXXMAMXMMMMSMMMMMAMAXXMAMAMXMMMSMMXMXMAXXAXAXMSAMXMXSAMAXAMXMSAXMMMSAXMAMAASMMSMMSMMSMSASXMAXMXMX
MMSSMASAMAMAMSXSAMMMMSASMSSMMMSXMAXXXMAMSASAMSSMMAMXAXMAXAAMSXSMSXMXAXAXSXMASASMMSMSMMXMASMSAMASASXMASAMSAMXAAMMXMASXMMAAXAAXAXMAMMMAAAMSMSM
SXAAMAMAMAMAMASMASMAASXMAMMXMXMXMAXMXMAMAMMAMXAASAMSMSSSMMSAMXMMSASMSMSSXASAMAMAAXMXAXXAMSXSASMSAMXMASAXSMXSMSXSASXSAASMMSMMSMMMSMAXSMMSXAXA
AMSSMSSMXSSXSSMSAMMMMMAMMMSASAAXSAMMMMASMMSXMSSMMXASAMAXXAMXSMMASAMAMAXMMMMMSSSMMSASMMMSAMASAMMMAMXMASMMXAASAMASAMASMMMMXXAXMAAAMMXMAAXSMMMS
MAMAMXAAAMXMXAAMXXXMMMXMXXSASXSXMASAAMXXXMMSAAAXAMXMAMMMMXSMSAMAMXMAMSMSXXAMAMXMAXMAAAAAAMXMMMXSXSMMASMMMMMMAMMMSMMMXSASXSSSMMMMXAXXMMMSAMXX
XXSAMSMMXSAAXMMMSMMSAMAXXMMMMMMAXASMSSMSSMAMMSSMAAXSXMXMXMAASMMASASXMAAAXSXMASMSSMSSSMXMXMXMAMMMMMAMASAXXAMSSMMAXXAMMMAXAMAAAXMAMSMSAAXSAMAS
SMSAMXMAXMAMXXSAAAASASXMASAAAAXXMASXAAAAXMASXXXAMSMMXSAXAAXXSASASMXMSMSMMMMSXMXAAAAXAXXSSSMSASAXASXMXSXSSMXAXAMXSXMSAMSMSMSSMMMASAAMSMMMMMAS
MASXMMMSSMXSAAMXSMMMAMXASMSSSSMSMXXXMMMMXSAMXSSSMMMMASASXMSASAMXMXAAAAXXAAAAMXMMMMMSXMAXAAXSASXSXSXXXMXMASMASXMXMXXMXMXAXAAAMXSXSMSMMMSMMMXM
SMSXXSAAXMAMMSMMXAMMAMMMXAAMAAAAMSMSXASAMMMSAMXMAAAMMMAMAAMMMAMAMSSSMMMSSSSMMXMAAAAMASMMSMMMAMASMMMMAMAMMMXXAMXAMASXSSMMMMSSMASAXXXAXAAMSMSX
MAMAMMMMSMMSMMAXSAMSSMSAMMMSXMSMXAAMXXMXMAAMASAMMSMSAMXMXMMAMAMXXAAAXSAMXMAAXAXSMMASMMSAAXAMMMMMAASMMXXXSAMXXSSXSAXMAAAXAXAAMXSMMMMSMSMSAAXM
MSMSMASXMAMXASXMMSMSAMMMSAAXXXAMMMXMMMSMMSXSAMASXMASMSMMMMSXSMMSXMSMMMMSSXSMMMMMASXMAAMSMMSXMASXSMSAXMSMMMXMXAAAMMSSSMMXMSXMAMMASXAXXMSSMMMA
AXAXAASMSAMXMMSMAMXSSMAXSMSSSSXSAMAASAMMAXAMXSAMAMAMAAXXAAAAAAAAAMMMXXMAXXAXASAMXSAMMMMASMXMSASXAASAMXXAAAASMMAMMSAMXXXXXAMXMASAMMSSXMAMAASM
SMXMMMSXAXXXAAXMASXMAMSXXAMAMAASAMSAMASMMXAMASXMAMSMSMSSMSXMSMMSXMASXSMASMMMMMMMAMAMSMSASXXAMMSXSMSASAMSMSXXAXMASMMSSMMXAASXXMMMSAAMMMXMMMSA
MAMXAMXMMASXMXSMSSMSAMXAMSMSMMMMXMXASAMASMMMAMASASXAAAXXAMMAXXAXAMXMMAMASAXAXAMMSSMMMXMXMXSSMXXAMXSXMAMMAXASMMMSXMAMAAMSSMMXAXAAMMXMASMMXAMX
SMAXXSAMXMSAMASMAXAMASMAMXAMXAXMAXXAMAMAMXAMSSMMSMMXMMMSXXASMMSSXMAMSXMAMMMSSMSAMAXMSAMASASAMXMXMAXASMMSAXXAASXMASMSXMMAMSSSMMSXSAMSMSAASMMM
SXSMAXMMAMXAMAXMXMMSAMXMSMSMSSMSAMMSMMSAMXMXMAMMAXAMMSMAXSMMAAAMASMMMMMMSXAXAAMASMMXSASXMXSXMXAAMXMXMSAMXMMSMMASAMAMMMMMSAXAAAXAXMMAXMMMMAAX
SAMMSXXXASXSMSSSMMXMXSXMAAAAMAXMASXXAMXAMAMASMMSMSMASAMXMMASMMXSAMMAMXMXAMXXMMSMMAXASAMXMAMXXASXSSXSAMXSAAXAXMAMXMMMAMSXMMSSMMMMMMSMSSXXSSMS
SAXAXAMSMMAXAMAXAMAMXMASMSMSMXMSSMXMMMSXXXSASXAAMAMXMXXSASAMAMXMASMXMXMAAAMMSASXSSMMMXMAMMSSMAMXAASMXSXSXSMXXMXSAMXSXSMAAXXAMXAXAXAXSMAXMAAA
MAMAMSMAAMAMAMAMMMAMASXMAXXXMMMXAAXXSAMXSXMXSMSSSMMSMSMMAMAMAMASAMMMMAMMXSAAMAMXAAAMAMSAMSAMXAMXSMSSMSAMAMASMMMMASAXMAXXMSXMASASMSMMSAMMSMMM
SXSXMAXMMMXSAMASXXXXAMXMAMMMSAMSMMMMMASXAAMASAAMAMMSAAXMSMMSASMMMSMASAMXAMMMMAMSSMMMAXMAMMAXSXSXXXMAMXAMXMAAAAASMMASXSSMMXMMXAAXAAXASAMXXASX
AMMAXASMSXMMAMAXAASMSMSMMXSASXMXMMSXSAMXMMXAMMMXSMAMXMAMAAXAAXXAMAXMXAXAAMXXXAXMAXMMXMSAXSXMMMMMMMSMMMMMXMMSSMXAMXXXXAAXXAXXAMSMMMSXSXMASMMM
SMMXMXXMAASMSMSSMXMAMAXXMAMMMMXASASXMAMMXMMMXMSAXMSSMXMSSSMMMMSAMXSXSAMSXMXMSSXSMMMSAXSMMSMAMASAAMAXMAXAMXAAXMASXMSMMSMMMMSMMSMXAAMMMMMXXMAS
MAASMMSMSMMAMAMXXASMMMMXMASXSXSAMASMSSMAAMASAAMASXAXMAMMAAAXXAXXASAMMAMMASXMAMAAXAAMMMMXAXAXSASXSMMSMSMAMMMXSMMXAAAXAAMAAAAXXAMXMSXMAAXASMXX
MMMSAAXAMMMAMAMMMMASMMSSSMSXAXMSMXMMAMMSASASMSMMMMMSXSSMMSMMMXSMMMASMXSMMMAMASMMMMMSXSSMSXSAMXSAMAXSAMSAMXAXSAMSMMMMSXSSMMAXMAMXXMAMSMSASMMS
XSXSMMMXMXMAMAMAAXMAXAMXAAMMMAAAXMAMAMAXMAMMXAMMMSXAMMAMAAAASMXAXSSMMASAMSXMASAAXMXSAMXAAAMSMMMAMMAMAMSXMMSXSAMMASXAMAXASXMXSAXSSSSMAAMMMMAA
MSAXMAAASXSXSMSSSSXSAMSXMMMAXSMMMSASXSSSSSSSXXXAAMMSXSAMXSAMXMASAMXAMASXMXMMASAMXMAMXMASMMMAXXMAMXMSAMXSMXXASXMXAMXSMXMAMXAAMMSXAAXMMMMSAMSS
AMXMMMMMMAMXAAMAAMAMMASXSASXMXAXXSASAAXAAAXMASMMMSAMMMAMAAXMASAXASMXMAMAMXAMXSXMAMASAMAXXMXASXXAXXASASAAMAMMMMSMMMXXXMMXMMMMSMXMMMMMXSXMAXXX
XMMSXSXSMMMSMMMMXMAMXMMASASMMSXMAMMMMMMMMMMMMMAMMMMSASXMASMSMSMSMMMAMMSMMXMSASAMASMSAMASASXASMSSXMMSMMMMMAXAAASASAXMMMXASASAMMXMXMAMASASMMMS
XSAMXMAXAAAAAMXSMMAMAXMAMSMAXMAMXSSMMXSAMXAXASXMXMASAMXMSMMXAXMAMAMXXAAXSMMMASAMASXSXMASAMMAMAAAASXSAAXXSXSMXMXAMMMSAAMSSXMASMSMASAMAMMMAXAM
XMASMMAMMMSSSMXMASXSSSMSSMSMAMAMXAAAAASASMSSMSAMAMMMAXXMAAMMAMSASXMXXXSXMAXMXMAMASXMASMMXMSSMMMSMMAMMXSAXXAXSXMMMXAXSXSXMASXMAAXAXMMAMXMMMSS
XSAMASXSXAXXMXMXAXAAAXXMSAASMSXXXMSMMXXXMAAMXSASXSMSMMSSSSMMMMMASMSMSXXASXMMMSMMXMASAMXAAXAAXMAXAMXMAXMAMXXMAMMSAMXMMXAXMASAMXMMSSSMMMSAXMXX
MMASXMAXMMMXMAAMXSMMMMSMMMMSMAMSMXAMXMSMMMMMAMMMMXAAXXXAMAMXAAMMXAAAXASMMAAAAAXXASAMAAMXMMSAMXXSXMMMMSMXMSSMXSASASXAAXMMMAMXXSAXAAAAAAMAMXXX
XSMMMMSMMAAAMMSAXAXAXAXAAXXXMAAAAXMSAMAAAXXMAXMAMMSSSSMXMASMSSSMMSMMMMMXMMXMMXSAMMASMMXSXMASMSXMAAMAAAXXMXAAAMXSAMMMAMXSAMXAXXAMXSSMMMSASMSM
MSAMXAAASMSSSMXAMMSMMXSMMXSASXSSSMASASMSMSSMAXSASAAAMAMXMASXAMXMAMAXAXAAMSSSXMMAXAASASASXSAMMSASXMSMSSMSMSMMXXAMXMASXMAXMSMXSSXXAMAXAASAMAAA
AMAXMAMMMMAAXAXXAXAMMMXAAAMMMAAAAMXMAMMMAMXMAMXAMMSMSAMXXAMMXMXMAMXSMSMXSAAXMASXMSXSAMASAMAXAMAMAAXAXMAAMXXXSAMXXAAMAMMXXMAMMAMXMMMMAMMMMSMX
MSMMXSSSMMXSMMMMSSXSXMMMMXSMMSMSMMAMAMXMAMXMASMAMMMXSASXMASAASXSMMXSAAXSMMSMXAAXAMAMXMAMMMSMMMAMMMMSMMSMSXMASXMMMMMXSASMMMSMSASASAXSSXMAMASA
MMASAMXMASXMAXAAMAMMASXMMXAMXXAAXSAMXMASAMXSAXXAMXAXXMAXSSSMMSAAASASMMMXASXAMMSMAMXMSXMXSAMAXSASXSMMAMXXMAMXXXXAAAAXXAAAAAAXSASASMMAAMXXSASX
MSAMXMASXMASAMMSSMASAXMASMMMMMSMMMASAMXSAMAXXMMMMMXMAXMAMXMAXMXMMMASXAXSXMXXAMAAXSSSMASMMMSAMSASAAAMMMSMMSXMASMSSXSMMXMSMSSXMAMAMMSMMSSMMXMX
AAXMMMMSASAAASAAAXAMMMMAMXMAXMAMMXXMAMASAMXXSMSASXMXXMSMMSSMMMMXXMMMMXXSAMXMMSASAMXASAMAAXMMMMAMMMSXMASAAAMMMAMAMXMXSAXAXAXXMAMXSAMXXAAASMMX
MSMMXSASAMMSMMMSSMMSXXMXSASMMSASAAMMXMXSAMXAMASXSAMSSMSAAAAAAMAMSMSASMMSAMXSXAMXMXSMMMSSXMASXSMSXAMAMASMMXXAMXMXMAMASASXSMMXXXXAMASAMXSSMAAX
AAAXASAMXMXAXAXXXASAMSMMSASAASMMMXSAXSMSXMMXXAMMSAMXAASMMSSSMMXAAAMASXAXMMAXMAMAMMSXMMMMXASXMAASMSXXMMXXMASMSMMMSMMMSAMMAMASXMSMSXMASXMAMSMS
SSSMXSASMSSSSMMSAXMAMXAAMXMMMSXMAAMMMXAXAXMMMMSAMAMXMMMMXXAAMXMSMSMAMMMSMMMSXSMASAMAXMAMXMAAXMMMMAMSXSSSMXSAAXMASAAASAMXAMXXAAAASASXSMXSXMAS
MMAMXSAMAAAAAAAAMMSMMSMMSASXAXMMMMSMAXSMMAXAAXMXSSMSXSASXSMMMMMAAXMSAXMAMASXAMSAMASXMMMSSMSMMMXXMAMXAXAAXAMXASXSSSMXMAMSXSXSMMMMSAMAMAAMAMAM
SXSMMMMMMMMMMMMSMAXAMXAASXSMSSXAASAMAMXAXSSXSASMXMAAAMAMAAAAAXMXSXXAMMSSSMSMMMMXMAMMSAMAXMAXAMXMSMSMMMSMMMSAMXMAMAMASXMMMMMAXAAMMMMAMSMXAMXS
AAMXXAASAMXAXAAXMSSMMSMMXMMXMAXSMSASMSSXMXAXMASMMMSMMMAMXMMMXMMSAMXMAXAMAMXAXAMMXASXSXMXSXSSSSMXAAAAMXMXXAAXXAMMMMSAXXAAAXMMSSMSASMSXMMSMMAA
MSMASMMSXMSMSMSSMAXMAXMXMASAXMAMXSAMXAMMMMMMMAMAXMXMXSSSXSAMSSMAAMMXXMXSAMXSMAASXXMAXAMAMMAMXAXXSSSSMAMXMASMMMMSAAMXSSSMXMAAAAXSASXMAXXAXMXM
XAXAAAAXAMAMAMXAXAMMSMAASXMXSMXXMXXSMMMMAAAMMMSMMXAMMMAAXMAXAAXXMMXSMMXSXSAASXMAMMMXMASAMASASMMXMAMXXAXAXMAXAXASMSSMMXAAAMMMMSMMMMASAMSMSSSS
SMSMSMMMXSASMSSSMXSAAASMSAMXMASMMMMXAXMMSSMMAMAMMSMXAAMMMSSMSSMSXSXMAXAXXSAMSSSMSMAASXMASAXMMASXMMASXMXSSMXSSMXSAAMAASMMMMMAAXAAASAMAXXAMAAA
SXXAMXSXAXAMXMAXXMMMSMXXSAMSAMXAAAXMSMMMAXMSMSASXMMSSSXXXAAMXAAXXSXXSMMSAMAXXAMAAMSMSAXXMAXASMMXSAAXXAXXAMAAMXMMXMAMMXAMAMSAMXSSMSXMMMMSMMMM
MXSXMAMMSSMXMXXMMMAAMAXAMAMXAMSMSXXAAAXMAMXAXXXXAAAAXXXMMMMSSMMMMMSAXAMAMSMMMMMSMXMXXMASMMSAMMXXXMAMMSMSAMMSMMMAMXXMSSSMAXAAXMAXAXMAAAXXAMXX
MAMMMXSAASAASMSAAMASXMMSMSMSMMSXAAMSMSMMAMSSSMASXMMSSMSMMSAXMASXMAMXXMMMXAAAAXMXXAXAMXXXAXMMMXMMMMSMAAAMSMXMASMAMSMXAAXMMMMAMXXMXMASMSSSSMSM
MXSASXMMSXSXSAASXMAXAMXAAAAXMAMMMMMXAXXXAXXAAMAMXAXAAAXAAMASMXMXMASMSSMXSSSMSSMAMSSMSMSMMMMASAAMSAAMSMSMAXMMAMXAMXMMMSMSMAXSXMASASXXMAXAAAAA
SSMXSAAXAMMMMXMMMMMSAASMSMSMMMXSXXSMMMSSMXMXMMSSSXMXMXMMMSSMMAMXSASAAAXXMAXXXAMAMXAMAXAAAMMAMASXMSXXAAXXMSASMSSMSAMXMAMAMSMXAXASASXMMSMMMSMS
MAMASXMMSAAAMAXAXAXSMMXMAMAAMMASMMMASAMASXAASMMMMAASMSAAMXAASMSMSMSMSMMSMMMMSMMMXXAMMSSXMSMXSMMAMXSSMSMMXXAAAAMMSMMASASASAXSXMMMAMXMAMAAAAXA
SMMASASMASMSXMSMMMXXSXXMXSMSMMASAXSAMXSAAMSMMAAAXMXAAXXMSSSMMXSASASXXAMXAXXAAXAAMSAMXAAASAXMAXSAMSXMMAAMSMSMMMSAMASASASXSMXSAMXMXMXMAXSMSMSX
AXMXMAXAMXXMXAAAMSMMAMSMMSAXXMASXMMXSMMMMXMASXMSXXMMMMSSXMMAXXMAMAMAMASMMMMSMSMXMAMXMMMMSASMMMMXMMAXXMMMAAAXAXMASMMMMAMXXMASAMASAMAMSXMXXASM
MMSSMMMSXSASMSMSMAAMMAAAAMAMXSXMXAMAXMASXMMMMXAAASAMXAXAAXSSMXMAMAMAMAMAAAXAAAASMMMXXXXXMXMMXAAAMMMMMSASMSMSMSSMMXXXMAMSSMMSXMASASMSMASAMASM
XAAAAXMAAMAMAXMAMXMMXSSSMMSMAXASMSMSMSAMAASXSMXMASASMMSSSMAXASMMSMSMSAMSMMMXXMMMAAMXSAMSSMMSMSMXSAMAAMXMXAXAMMAAXSMMXSSXAXXXAMMSAMXAMXMAMASM
MMMXMMMMXMMMSMSMSSMSXMXAAXXMAXXAAAAAMMMSMMMAXXXAASAMAMXAMMXMMMMAAXAMMXMMAAASXSASMMMSMAMAAMXSAXAASAMMXMAASMSSSSMMMSAAAXMXMSMMAMXMMMSXMMXXMAMX
MSSXMAAXAAXXMXMXAAXSASMSASMSASMMSMSMSXAAAAMMMMSMMMMSSMMAMMMSXAXSMXMSASMSXMMSAMMXAAXXMAMSSMXMAMAMMXMAXXMMAMAAMAASASMMMSAMXSAXSXMASAMXMAXSMSMS
AAAAMSSXMSSMSAMXSMMSXMAXAAAAAAXXAXXAXMSSSMMXAAAXXAAMAMMSSXAASXMASXXMXSXMASXMAMXXXMMSSMXXAAAMXMAXSSMMMASXSMMSMSMMASAAXAAMASMMAASAMXXSXXXAAAMA
MMSSMAMMXMAXSXSAMAAMMMSMMMSMSMSSXXMXMXAXXXXSMXSMMMSSMMAMAASMAMSAMXXMASASASASXMMSAAXXAMMMMMMMAMMSMXASXMSAMSAAXAXMXMMSSSMMXSXXSAMMSAXSASAMSMSM
XXAXMAMMASMMXMMAMMXSAAAAXAMXXAAAMSMSAMXSMSMSAAXMXAMAMMMMSAMXMMXAMXMXASAMASXMXAASXMMSMMSAAAMXSXSAMMAMAXMAMMMMSXXXMMSAMXAXAXXXXXMAMMSMAMMXMAMX
MMMSMMMSASXAAASMMAXSMMSSMMSMMXMMMAAAMAMXAXAMMMMAMSMMMXSAMAXXXXMSAMXSXMXMAXMMMMMSASAAXASXXMSAMMSASMSSMMMSMSSMMMMAAXMASXMMXSMMMXMASXXMAMXMMAMM
AAMXMAMMAMXXMMAAXAAXAAXXAMXAAXMSSMMMASAMXMSMSASMMXAAAASXSSMSMSAMASASAMXMSSMXAAAMMMAMMASAXSMXSASXMMAAAXMAAAAAAXSASMSMMMMSAXAAAXSXMMMSXXXAMXMA
SXSASXMXASXSMMSMMXXSMMMSSMMSMMSAMXXXAMAMSAMASAMXAMSMMMSAXXAAASASAMASXMAAMAXASMSXXXMXSXMMMAASMMMSAMXSMMSMSMSMMMSAXAAAMXAMMMSSSMMAMMMAMMMSXASM
XASASAXSASASXAAAXMXMASXAXXAAASMMXSAMXSAMMSSMMSMMMMXMSAMXMMSMXMMMXMXMASMSMMMAXAMMSXMASXMAXMMMAXAXXSAMXAXAMXXAMXMAMSMSMMXSXAXMAMSAMAMASAAXSMMX
MAMAMAMMAMAMMSMSAMMSAMMXMMSSSMAAXMSAASASAASMAAXASAMXMASASAXSAASAXMASAMSAASMXMAMAMXMASASMSXMSAMXSAMMSMASXMASXMAMMMXMXAXAMMSXSAMXASXSASMSMSXAM
MSMSMAMMAMXMAMMXMAMXSSXAXXXXXXMMSAMXMSSMMSSMSSSXSXAXSMMMAXMASXMMXSAMXSMSSXAMSXMASAMXSAMXSAMXMAAAXSXXXMAMXMASXXSMXMASAMXXSAMXXSSMMXMXSMMAMXXS`
}
