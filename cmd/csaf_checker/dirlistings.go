// This file is Free Software under the MIT License
// without warranty, see README.md and LICENSES/MIT.txt for details.
//
// SPDX-License-Identifier: MIT
//
// SPDX-FileCopyrightText: 2021 German Federal Office for Information Security (BSI) <https://www.bsi.bund.de>
// Software-Engineering: 2021 Intevation GmbH <https://intevation.de>

package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func (p *processor) collectLinks(dir string, consume func(string)) error {

	client := p.httpClient()

	res, err := client.Get(dir)
	if err != nil {
		return nil, err
	}

	if res.Status != http.StatusOK {
		return fmt.Errorf("status %d (%s)", res.StatusCode, res.Status)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if link, ok := s.Attribute("href"); ok {
			consume(link)
		}
	})

	return links, nil
}
