## Version 0.4.0, 2026.01.06

* **Breaking Change:** Remove `Coder` type
* **Breaking Change:** Remove `Generator.Coder` method
* **Breaking Change:** Rename following fields:
    * `Generator.CodeNSValidator` to `Generator.CodeNamespaceValidator`
    * `ParsedCode.NS` to `ParsedCode.Namespace`
* **Breaking Change:** Rename following types:
  * `NS` to `CodeNamespace`
  * `NSValidator` to `CodeNamespaceValidator`
* **Breaking Change:** Rename following functions:
    * `ComposeNSValidator` to `ComposeCodeNamespaceValidator`
    * `LenNSValidator` to `LenCodeNamespaceValidator`
    * `RegexpNSValidator` to `RegexpCodeNamespaceValidator`
    * `UnicodeNSValidator` to `UnicodeCodeNamespaceValidator`
    * `HasCodeNS` to `HasCodeNamespace`
    * `HasCodeNSUsing` to `HasCodeNamespaceUsing`
* **Breaking Change:** Change signature of following functions:
  * `BuildCode(uint, ...NS) (Code, error)` to `BuildCode(uint, CodeNamespace) (Code, error)`
  * `HasCode(Code, ...Operator) Matcher` to `HasCode(uint, CodeNamespace, ...Operator) Matcher`
  * `MustBuildCode(uint, ...NS) Code` to `MustBuildCode(uint, CodeNamespace) Code`
  * `WithCode(Code) Option` to `WithCode(uint, CodeNamespace) Option`
* **Breaking Change:** Change signature of `Builder.Code(Code) *Builder` method to
  `Builder.Code(uint, CodeNamespace) *Builder`
* Add following functions for working with codes:
    * `HasCodeUsing`
    * `MustValidateCodeNamespace`
    * `MustValidateCodeValue`
    * `ValidateCodeNamespace`
    * `ValidateCodeValue`
* Add following methods for working with codes:
    * `Generator.BuildCode`
    * `Generator.MustBuildCode`
    * `Generator.MustParseCode`
    * `Generator.MustValidateCode`
    * `Generator.MustValidateCodeNamespace`
    * `Generator.MustValidateCodeValue`
    * `Generator.ParseCode`
    * `Generator.ValidateCode`
    * `Generator.ValidateCodeNamespace`
    * `Generator.ValidateCodeValue`
* Add `Problem.Clone` method to safely clone a `Problem`
* Add `V7UUIDGenerator` and `V7UUIDGeneratorFromReader` functions to return a `UUIDGenerator` that generates a (V7) UUID
* Bump all dependencies to latest versions

## Version 0.3.1, 2025.10.09

* Add `LICENSE.md` file for `github.com/neocotic/go-problem/zap` module

## Version 0.3.0, 2025.10.09

* **Breaking Change:** Remove built-in support for [zap](https://github.com/uber-go/zap) logging
* Add `github.com/neocotic/go-problem/zap` module to provide support for [zap](https://github.com/uber-go/zap) logging

## Version 0.2.1, 2025.09.30

* Fix link to `go get` command in `README.md`

## Version 0.2.0, 2025.09.30

* Bump support for Go to 1.24+
* Change all functions that return `Builder` to return `*Builder` pointer instead
* Add built-in support for [zap](https://github.com/uber-go/zap) logging
* Fix issue causing `Translator` to be ignored when specified
* Improve documentation
* Bump all dependencies to latest versions

## Version 0.1.0, 2024.05.13

* Initial release
