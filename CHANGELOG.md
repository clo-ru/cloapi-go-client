# Release v3.1.0 (2026-06-11)
* **Breaking.** The client is now generated from the CLO OpenAPI spec with
  oapi-codegen. The hand-written `services/*` request structs are removed; call the
  generated `*WithResponse` methods on the client returned by `cloapi.New` instead.
* Single root package `cloapi` (module path `github.com/clo-ru/cloapi-go-client/v3`).
* Hand-written ergonomic layer: functional-option client (`WithBaseURL`, `WithTimeout`,
  `WithLogger`, `WithHTTPClient`, `WithRetry`), logging + retry transports, error
  classification (`IsNotFound`/`IsClientError`/`IsServerError`), query/filter editors,
  and a generic `Paginator[T]`.
* Daily regeneration keeps the client in sync with the API.
* See [MIGRATION.md](MIGRATION.md) for the v2 → v3 upgrade guide. The `/v2` import path
  is frozen but remains available at its existing tags.

# Release v1.0.0 (2022-08-31)
* Initial release with a core functionality