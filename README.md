# CVEdash

CVEdash (alternative spelling: `CVE-`) is a dashboard intended to give an overview of the latest published CVEs and some nice statistics.

The project is still in early stages, with more features yet to come which will make the dashboad more useful. These features include search, filters, detail information and many more.

This is my first larger project using Go and SvelteKit, so the code may still be a little messy.

I chose a mono-repo approach to bundle backend ([api](api)) and frontend ([app](app)) in one single repository, since both projects are still quite small and I like to have things in one place.

## Disclaimer
This product uses the NVD API but is not endorsed or certified by the NVD.