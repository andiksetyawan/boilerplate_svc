# go install github.com/vektra/mockery/v2@v2.43.0

# internal
##  usecase
mockery --dir=internal/usecase --output=mocks/usecase --outpkg=usecasemock --all

##  repository
mockery --dir=internal/repository --output=mocks/repository --outpkg=repomock --all
