// PhishDetect
// Copyright (C) 2018  Claudio Guarnieri
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package brand

// WhatsApp brand properties.
func WhatsApp() *Brand {
	name := "whatsapp"
	original := []string{"whatsapp"}
	whitelist := []string{"whatsapp.com", "whatsapp.org", "whatsapp.net"}
	suspicious := []string{
		"whatsappa", "whatsappb", "whatsappc", "whatsappd", "whatsappe", "whatsappf", "whatsappg", "whatsapph", "whatsappi", "whatsappj", "whatsappk", "whatsappl", "whatsappm", "whatsappn", "whatsappo", "whatsappp", "whatsappq", "whatsappr", "whatsapps", "whatsappt", "whatsapp", "whatsappv", "whatsappw", "whatsappx", "whatsappy", "whatsappz", "vhatsapp", "uhatsapp", "shatsapp", "ghatsapp", "7hatsapp", "wiatsapp", "wjatsapp", "wlatsapp", "wxatsapp", "whctsapp", "whetsapp", "whitsapp", "whqtsapp", "whausapp", "whavsapp", "whapsapp", "whadsapp", "wha4sapp", "whatrapp", "whatqapp", "whatwapp", "whatcapp", "what3app", "whatscpp", "whatsepp", "whatsipp", "whatsqpp", "whatsaqp", "whatsarp", "whatsatp", "whatsaxp", "whatsa0p", "whatsapq", "whatsapr", "whatsapt", "whatsapx", "whatsap0", "xn--whatapp-gpb", "xn--whatspp-6fg", "xn--whtspp-4tac", "xn--whatsap-vjg", "xn--whatsap-f2f", "xn--whatsa-1cfa", "vvhatsapp", "xn--hatsapp-7yh", "xn--whtspp-4nfc", "xn--whtspp-xocc", "xn--whtspp-qtac", "xn--whatsap-92a", "xn--whtspp-cxcc", "xn--whtspp-45bc", "xn--whtsapp-6wa", "xn--whasapp-8jg", "xn--whasapp-8qb", "xn--whtspp-qlgc", "xn--hatsapp-rfi", "xn--whatsa-nzaa", "xn--whatspp-lwa", "xn--whtsapp-dih", "xn--whatspp-pbd", "xn--whtspp-xtac", "xn--whtspp-juac", "xn--whatspp-1wa", "xn--watsapp-7ii", "xn--whatspp-dwa", "xn--whtspp-xc8bc", "xn--whtsapp-93a", "xn--whtsapp-iwa", "xn--whtsapp-ogc", "xn--whatsap-ddc", "xn--whasapp-gqf", "xn--whatspp-twa", "xn--whatspp-en4c", "xn--whatsap-vpf", "xn--whatapp-wu8e", "wihatsapp", "xn--whtsapp-00c", "xn--watsapp-5g6a", "xn--whatapp-qm6a", "xn--whtsapp-mbd", "xn--whatspp-9wa", "xn--whatsap-upf", "xn--whatsap-82a", "xn--whatapp-vmd", "xn--watsapp-uch", "xn--whtsapp-ywa", "xn--hatsapp-8qg", "xn--whatsap-wjg", "xn--whatapp-mog", "wlhatsapp", "xn--whtspp-cuac", "xn--whtsapp-3fg", "xn--whatsa-grfa", "xn--whtspp-j0ac", "xn--whatspp-d4a", "xn--whatsa-82ba", "xn--whatspp-30c", "xn--whatsap-g2f", "xn--whtsapp-9va", "xn--whatsap-edc", "xn--whtsapp-fxa", "xn--whtsapp-qwa", "xn--whtspp-jtac", "xn--whtsapp-bn4c", "xn--watsapp-1ig", "xn--whatsa-10ea", "xn--whatspp-rgc", "xn--whatspp-ixa", "xn--whatspp-gih", "w-hatsapp", "wh-atsapp", "wha-tsapp", "what-sapp", "whats-app", "whatsa-pp", "whatsap-p", "whatsaplp", "wha2tsapp", "whats1app", "whatzsapp", "whatsampp", "whzatsapp", "whsatsapp", "whatsa1pp", "wha5tsapp", "whatesapp", "wzhatsapp", "whyatsapp", "what6sapp", "whatsaapp", "whnatsapp", "whatsap0p", "whartsapp", "whwatsapp", "wjhatsapp", "wghatsapp", "whatsqapp", "wyhatsapp", "wbhatsapp", "whaztsapp", "whatsyapp", "whatsa2pp", "whatsawpp", "whatrsapp", "whatszapp", "whatsapmp", "whatsaopp", "whatsxapp", "whjatsapp", "whqatsapp", "whats2app", "whatsaspp", "whatxsapp", "whatqsapp", "whatgsapp", "whatfsapp", "whagtsapp", "whatdsapp", "whatseapp", "whastsapp", "whbatsapp", "whatwsapp", "whawtsapp", "whatsa0pp", "whatsdapp", "whatasapp", "wha1tsapp", "whuatsapp", "whatssapp", "what5sapp", "whaytsapp", "whatysapp", "whatsaqpp", "wh1atsapp", "whatsalpp", "whatswapp", "whgatsapp", "wnhatsapp", "wh2atsapp", "whatsapop", "whaftsapp", "wha6tsapp", "wuhatsapp", "whaqtsapp", "whatsaypp", "whatsazpp", "hatsapp", "whatapp", "whatspp", "whtsapp", "whatsap", "whasapp", "watsapp", "whhatsapp", "whaatsapp", "wwhatsapp", "whattsapp", "wha5sapp", "wyatsapp", "whstsapp", "whats1pp", "whats2pp", "xhatsapp", "whatsamp", "wuatsapp", "ahatsapp", "whatsapo", "whatsapl", "wh2tsapp", "wbatsapp", "whatzapp", "whytsapp", "whatdapp", "wha6sapp", "2hatsapp", "whatswpp", "3hatsapp", "whaysapp", "whatsapm", "whztsapp", "wharsapp", "wgatsapp", "whatsaop", "qhatsapp", "whatsspp", "whafsapp", "whatxapp", "whataapp", "whatsypp", "wnatsapp", "whagsapp", "whatsalp", "ehatsapp", "whatszpp", "wh1tsapp", "whatyapp", "whateapp", "wzatsapp", "whazsapp", "whwtsapp", "w.hatsapp", "wh.atsapp", "wha.tsapp", "what.sapp", "whats.app", "whatsa.pp", "whatsap.p", "hwatsapp", "wahtsapp", "whtasapp", "whastapp", "whataspp", "whatspap", "whatsopp", "whutsapp", "whotsapp", "whatsupp", "whatsappcom",
	}

	return &Brand{
		Name:       name,
		Original:   original,
		Whitelist:  whitelist,
		Suspicious: suspicious,
	}
}
