package index

import "github.com/intdxdt/mbr"

func (tree *Index) Collides(query mbr.MBR) bool {
    var bbox = &query
    if !intersects(bbox, &tree.data.bbox) {
        return false
    }
    var child *idxNode
    var bln  = false
    var searchList []*idxNode
    var nd = &tree.data

    for !bln && nd != nil {
        for i, length := 0, len(nd.children); !bln && i < length; i++ {
            child = &nd.children[i]
            if intersects(bbox, &child.bbox) {
                bln =  nd.leaf || contains(bbox, &child.bbox)
                searchList = append(searchList, child)
            }
        }
        nd, searchList = popNode(searchList)
    }
    return bln
}
