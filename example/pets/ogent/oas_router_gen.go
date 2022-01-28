// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "DELETE":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleDeletePetRequest([1]string{
					args[0],
				}, w, r)

				return
			}
			switch elem[0] {
			case 'c': // Prefix: "categories/"
				if l := len("categories/"); len(elem) >= l && elem[0:l] == "categories/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteCategory
					s.handleDeleteCategoryRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			case 'p': // Prefix: "pets/"
				if l := len("pets/"); len(elem) >= l && elem[0:l] == "pets/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx > 0 {
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						s.handleDeletePetRequest([1]string{
							args[0],
						}, w, r)

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/owner"
						if l := len("/owner"); len(elem) >= l && elem[0:l] == "/owner" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: DeletePetOwner
							s.handleDeletePetOwnerRequest([1]string{
								args[0],
							}, w, r)

							return
						}
					}
				}
			case 'u': // Prefix: "users/"
				if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteUser
					s.handleDeleteUserRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			}
		}
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleListPetRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case 'c': // Prefix: "categories"
				if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleListCategoryRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							s.handleReadCategoryRequest([1]string{
								args[0],
							}, w, r)

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/pets"
							if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ListCategoryPets
								s.handleListCategoryPetsRequest([1]string{
									args[0],
								}, w, r)

								return
							}
						}
					}
				}
			case 'p': // Prefix: "pets"
				if l := len("pets"); len(elem) >= l && elem[0:l] == "pets" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleListPetRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							s.handleReadPetRequest([1]string{
								args[0],
							}, w, r)

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleListPetFriendsRequest([1]string{
									args[0],
								}, w, r)

								return
							}
							switch elem[0] {
							case 'c': // Prefix: "categories"
								if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: ListPetCategories
									s.handleListPetCategoriesRequest([1]string{
										args[0],
									}, w, r)

									return
								}
							case 'f': // Prefix: "friends"
								if l := len("friends"); len(elem) >= l && elem[0:l] == "friends" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: ListPetFriends
									s.handleListPetFriendsRequest([1]string{
										args[0],
									}, w, r)

									return
								}
							case 'o': // Prefix: "owner"
								if l := len("owner"); len(elem) >= l && elem[0:l] == "owner" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: ReadPetOwner
									s.handleReadPetOwnerRequest([1]string{
										args[0],
									}, w, r)

									return
								}
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleListUserRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							s.handleReadUserRequest([1]string{
								args[0],
							}, w, r)

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/pets"
							if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ListUserPets
								s.handleListUserPetsRequest([1]string{
									args[0],
								}, w, r)

								return
							}
						}
					}
				}
			}
		}
	case "PATCH":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleUpdatePetRequest([1]string{
					args[0],
				}, w, r)

				return
			}
			switch elem[0] {
			case 'c': // Prefix: "categories/"
				if l := len("categories/"); len(elem) >= l && elem[0:l] == "categories/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateCategory
					s.handleUpdateCategoryRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			case 'p': // Prefix: "pets/"
				if l := len("pets/"); len(elem) >= l && elem[0:l] == "pets/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdatePet
					s.handleUpdatePetRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			case 'u': // Prefix: "users/"
				if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateUser
					s.handleUpdateUserRequest([1]string{
						args[0],
					}, w, r)

					return
				}
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleCreatePetRequest([0]string{}, w, r)

				return
			}
			switch elem[0] {
			case 'c': // Prefix: "categories"
				if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleCreateCategoryRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/pets"
							if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: CreateCategoryPets
								s.handleCreateCategoryPetsRequest([1]string{
									args[0],
								}, w, r)

								return
							}
						}
					}
				}
			case 'p': // Prefix: "pets"
				if l := len("pets"); len(elem) >= l && elem[0:l] == "pets" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleCreatePetRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleCreatePetFriendsRequest([1]string{
									args[0],
								}, w, r)

								return
							}
							switch elem[0] {
							case 'c': // Prefix: "categories"
								if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: CreatePetCategories
									s.handleCreatePetCategoriesRequest([1]string{
										args[0],
									}, w, r)

									return
								}
							case 'f': // Prefix: "friends"
								if l := len("friends"); len(elem) >= l && elem[0:l] == "friends" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: CreatePetFriends
									s.handleCreatePetFriendsRequest([1]string{
										args[0],
									}, w, r)

									return
								}
							case 'o': // Prefix: "owner"
								if l := len("owner"); len(elem) >= l && elem[0:l] == "owner" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: CreatePetOwner
									s.handleCreatePetOwnerRequest([1]string{
										args[0],
									}, w, r)

									return
								}
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleCreateUserRequest([0]string{}, w, r)

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/pets"
							if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: CreateUserPets
								s.handleCreateUserPetsRequest([1]string{
									args[0],
								}, w, r)

								return
							}
						}
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name  string
	count int
	args  [1]string
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.name
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [1]string{}
		elem = path
	)
	r.args = args

	// Static code generated router with unwrapped path search.
	switch method {
	case "DELETE":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "DeletePet"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'c': // Prefix: "categories/"
				if l := len("categories/"); len(elem) >= l && elem[0:l] == "categories/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteCategory
					r.name = "DeleteCategory"
					r.args = args
					r.count = 1
					return r, true
				}
			case 'p': // Prefix: "pets/"
				if l := len("pets/"); len(elem) >= l && elem[0:l] == "pets/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx > 0 {
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						r.name = "DeletePet"
						r.args = args
						r.count = 1
						return r, true
					}
					switch elem[0] {
					case '/': // Prefix: "/owner"
						if l := len("/owner"); len(elem) >= l && elem[0:l] == "/owner" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: DeletePetOwner
							r.name = "DeletePetOwner"
							r.args = args
							r.count = 1
							return r, true
						}
					}
				}
			case 'u': // Prefix: "users/"
				if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: DeleteUser
					r.name = "DeleteUser"
					r.args = args
					r.count = 1
					return r, true
				}
			}
		}
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "ListPet"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'c': // Prefix: "categories"
				if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "ListCategory"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							r.name = "ReadCategory"
							r.args = args
							r.count = 1
							return r, true
						}
						switch elem[0] {
						case '/': // Prefix: "/pets"
							if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ListCategoryPets
								r.name = "ListCategoryPets"
								r.args = args
								r.count = 1
								return r, true
							}
						}
					}
				}
			case 'p': // Prefix: "pets"
				if l := len("pets"); len(elem) >= l && elem[0:l] == "pets" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "ListPet"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							r.name = "ReadPet"
							r.args = args
							r.count = 1
							return r, true
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								r.name = "ListPetFriends"
								r.args = args
								r.count = 1
								return r, true
							}
							switch elem[0] {
							case 'c': // Prefix: "categories"
								if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: ListPetCategories
									r.name = "ListPetCategories"
									r.args = args
									r.count = 1
									return r, true
								}
							case 'f': // Prefix: "friends"
								if l := len("friends"); len(elem) >= l && elem[0:l] == "friends" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: ListPetFriends
									r.name = "ListPetFriends"
									r.args = args
									r.count = 1
									return r, true
								}
							case 'o': // Prefix: "owner"
								if l := len("owner"); len(elem) >= l && elem[0:l] == "owner" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: ReadPetOwner
									r.name = "ReadPetOwner"
									r.args = args
									r.count = 1
									return r, true
								}
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "ListUser"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							r.name = "ReadUser"
							r.args = args
							r.count = 1
							return r, true
						}
						switch elem[0] {
						case '/': // Prefix: "/pets"
							if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: ListUserPets
								r.name = "ListUserPets"
								r.args = args
								r.count = 1
								return r, true
							}
						}
					}
				}
			}
		}
	case "PATCH":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "UpdatePet"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'c': // Prefix: "categories/"
				if l := len("categories/"); len(elem) >= l && elem[0:l] == "categories/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateCategory
					r.name = "UpdateCategory"
					r.args = args
					r.count = 1
					return r, true
				}
			case 'p': // Prefix: "pets/"
				if l := len("pets/"); len(elem) >= l && elem[0:l] == "pets/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdatePet
					r.name = "UpdatePet"
					r.args = args
					r.count = 1
					return r, true
				}
			case 'u': // Prefix: "users/"
				if l := len("users/"); len(elem) >= l && elem[0:l] == "users/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf: UpdateUser
					r.name = "UpdateUser"
					r.args = args
					r.count = 1
					return r, true
				}
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				r.name = "CreatePet"
				r.args = args
				r.count = 0
				return r, true
			}
			switch elem[0] {
			case 'c': // Prefix: "categories"
				if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "CreateCategory"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/pets"
							if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: CreateCategoryPets
								r.name = "CreateCategoryPets"
								r.args = args
								r.count = 1
								return r, true
							}
						}
					}
				}
			case 'p': // Prefix: "pets"
				if l := len("pets"); len(elem) >= l && elem[0:l] == "pets" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "CreatePet"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"
							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								r.name = "CreatePetFriends"
								r.args = args
								r.count = 1
								return r, true
							}
							switch elem[0] {
							case 'c': // Prefix: "categories"
								if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: CreatePetCategories
									r.name = "CreatePetCategories"
									r.args = args
									r.count = 1
									return r, true
								}
							case 'f': // Prefix: "friends"
								if l := len("friends"); len(elem) >= l && elem[0:l] == "friends" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: CreatePetFriends
									r.name = "CreatePetFriends"
									r.args = args
									r.count = 1
									return r, true
								}
							case 'o': // Prefix: "owner"
								if l := len("owner"); len(elem) >= l && elem[0:l] == "owner" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: CreatePetOwner
									r.name = "CreatePetOwner"
									r.args = args
									r.count = 1
									return r, true
								}
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					r.name = "CreateUser"
					r.args = args
					r.count = 0
					return r, true
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx > 0 {
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/pets"
							if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: CreateUserPets
								r.name = "CreateUserPets"
								r.args = args
								r.count = 1
								return r, true
							}
						}
					}
				}
			}
		}
	}
	return r, false
}
