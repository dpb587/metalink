package cmd

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/dpb587/blob-receipt/repository/filter"
	filter_and "github.com/dpb587/blob-receipt/repository/filter/and"
	"github.com/dpb587/blob-receipt/repository/sorter"
	sorter_reverse "github.com/dpb587/blob-receipt/repository/sorter/reverse"
	sorter_v "github.com/dpb587/blob-receipt/repository/sorter/v"
	// sorter_reverse "github.com/dpb587/blob-receipt/repository/sorter/reverse"
	// sorter_v "github.com/dpb587/blob-receipt/repository/sorter/v"
	// "github.com/dpb587/blob-receipt/repository/sorter"
	"github.com/dpb587/blob-receipt/repository/source"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type List struct {
	Filter []string `short:"f" long:"filter" description:"Filter blob receipts" default-mask:"TYPE:KEY[:VALUE]"`
	Sort   string   `short:"s" long:"sort" description:"Sort blob receipts" default-mask:"TYPE:KEY[:ORDER]"`
	Limit  int      `short:"n" long:"limit" description:"Limit the number of blob receipts"`
	Raw    bool     `long:"raw" description:"Output full receipt JSON"`
	Args   ListArgs `positional-args:"true" required:"true"`

	SourceFactory source.Factory
	FilterManager filter.Manager
}

type ListArgs struct {
	RepositoryURI string `positional-arg-name:"URI" description:"Repository URI hosting the receipts"`
}

func (c *List) Execute(_ []string) error {
	repository, err := c.SourceFactory.Create(c.Args.RepositoryURI)
	if err != nil {
		return bosherr.WrapError(err, "Creating repository")
	}

	err = repository.Reload()
	if err != nil {
		return bosherr.WrapError(err, "Loading repository")
	}

	andFilter := filter_and.NewFilter()

	for _, filterArg := range c.Filter {
		// @todo goflag arg
		filterArgSplit := strings.SplitN(filterArg, ":", 3)
		if len(filterArgSplit) != 3 {
			panic("unexpected filter arg")
		}

		addFilter, err := c.FilterManager.CreateFilter([]map[string]interface{}{
			map[string]interface{}{
				filterArgSplit[0]: map[string]string{
					filterArgSplit[1]: filterArgSplit[2],
				},
			},
		})
		if err != nil {
			panic(err)
		}

		andFilter.Add(addFilter)
	}

	receipts, err := repository.FilterBlobReceipts(andFilter)
	if err != nil {
		return bosherr.WrapError(err, "Filtering blob receipts")
	}

	var sort sorter.Sorter

	if c.Sort != "" {
		sorterArgSplit := strings.SplitN(c.Sort, ":", 3)
		if len(sorterArgSplit) < 2 {
			panic("unexpected sorter arg")
		}

		sort = sorter_v.Sorter{
			Field: sorterArgSplit[1],
		}

		if len(sorterArgSplit) == 3 && strings.ToLower(sorterArgSplit[2]) == "asc" {
			sort = sorter_reverse.Sorter{
				Sorter: sort,
			}
		}

		sorter.Sort(receipts, sort)
	}

	limit := c.Limit
	if limit < 0 {
		panic("Invalid limit")
	} else if limit == 0 {
		limit = 10
	}

	limit = int(math.Min(float64(len(receipts)), float64(limit)))

	for _, receipt := range receipts[:limit] {
		if c.Raw {
			marshal, err := json.MarshalIndent(receipt, "", "  ")
			if err != nil {
				// @todo
				fmt.Println(fmt.Sprintf("%#+v", err))
			}

			fmt.Println(fmt.Sprintf("%s", marshal))
		} else {
			fmt.Println(receipt.Repository.Path)
		}
	}

	return nil
}
