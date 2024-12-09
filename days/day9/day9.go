package day9

import (
	"fmt"
	"strconv"
	"adventofcode/m/v2/util"
)

func Day9(inputFile string, part int) {
	if part == 0 {
		fmt.Printf("Compressed file checksum: %d\n", compress(inputFile))
	} else {
		fmt.Printf("Compressed file checksum: %d\n", compressIntact(inputFile))
	}
}

func compress(inputFile string) int {
	ls := util.LineScanner(inputFile)
	line, _ := util.Read(ls)

	var id, idx, blocks, checksum int 
	j := len(line)-1 // Original file system: Index of the last file
	for i, r := range line {
		if j < i { // Stop when we're past the last file
			break
		}
		if i % 2 == 1 { // Odd positions: free space
			k := 0 
			for k < int(r)-48 { // for every free space we move a file block
				if blocks == 0 { 
					blocks = int(line[j])-48 
				}
				if blocks > 0 {
					id = j/2
					checksum += idx*id
					idx++
					blocks--
				}
				if blocks == 0 {
					j -= 2 // move index to the next file to move
				}
				if j < i { 
					break
				}
				k++
			}
		} else { // Even positions: file
			id = i/2 
			k := 0
			bb := int(r)-48
			if i == j { // Remember to get the current block value if we're at the last file
				bb = blocks
			} 
			for k < bb { 
				checksum += idx*id
				idx++
				k++
			}
		}
	}

	return checksum
}

func compressIntact(inputFile string) int {
  ls := util.LineScanner(inputFile)
	line, _ := util.Read(ls)

	// indices[i] = index at the start of file block i or space block i
	var idx, id, checksum int
	indices := []int{}
	for _, r := range(line) {
		indices = append(indices, idx)
		idx += int(r)-48
	}

	for j := len(line)-1; j > 0; j-=2 {
		blocks := int(line[j])-48
		// Search for space for these blocks
		for i := 1; i < j; i+=2 {
			if blocks <= int(line[i])-48 {
				for k := 0; k < blocks; k++ {
					// Move the files and add to checksum
					idx = indices[i]+k
					id = j/2
					checksum += idx*id
				}
				
				indices[i] += blocks
				
				// make sure file is moved and the free space is adjusted
				sizeLeft := int(line[i])-48-blocks
				line = line[:i]+strconv.Itoa(sizeLeft)+line[i+1:]
				line = line[:j-1]+"xx"+line[j+1:] // A file is only moved once so we can ignore the space ahead of it

				break
			}
		}
	}

	// Add the files that didn't move to checksum
	id = 0
	for i := 0; i < len(line); i += 2 {
		if line[i] == 'x' {
			id++
			continue
		}

		blocks := int(line[i])-48
		for k := 0; k < blocks; k++ {
			idx = indices[i]+k
			//fmt.Printf("Adding to checksum idx(%d) * id(%d)\n", idx, id)
			checksum += idx*id
		}
		id++
	}	
	return checksum
}
