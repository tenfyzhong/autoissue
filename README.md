# autoissue
hexo创建博客后，自动初始化gitalk/gitment评论


## Build

Use Go 1.26.10 (or Go 1.27.3 or newer on that release line), which includes
the `crypto/x509` fix for [GO-2026-6809](https://pkg.go.dev/vuln/GO-2026-6809).
This tool sends HTTPS requests to the GitHub API, so rebuild existing binaries
to pick up the fixed standard library.

```sh
git clone https://github.com/tenfyzhong/autoissue.git
cd autoissue
GOTOOLCHAIN=go1.26.10 go install .
```

The Go module preserves the dependency revisions from `vendor.conf` and
`trash.lock`. The toolchain directive selects the patched compiler with modern
Go tooling without moving the source to a new language version.
CI pins and checks the actual compiler and built executable; it does not create
issues, publish releases, or deploy the application.
