// Code generated by ent, DO NOT EDIT.

package ogent

import "github.com/tiagoposse/ogent/internal/integration/ogent/ent"

func NewAllTypesCreate(e *ent.AllTypes) *AllTypesCreate {
	if e == nil {
		return nil
	}
	var ret AllTypesCreate
	ret.ID = int64(e.ID)
	ret.Int = e.Int
	ret.Int8 = int32(e.Int8)
	ret.Int16 = int32(e.Int16)
	ret.Int32 = e.Int32
	ret.Int64 = e.Int64
	ret.Uint = int64(e.Uint)
	ret.Uint8 = int32(e.Uint8)
	ret.Uint16 = int32(e.Uint16)
	ret.Uint32 = int64(e.Uint32)
	ret.Uint64 = int64(e.Uint64)
	ret.Float32 = e.Float32
	ret.Float64 = e.Float64
	ret.StringType = e.StringType
	ret.Bool = e.Bool
	ret.UUID = e.UUID
	ret.Time = e.Time
	ret.Text = e.Text
	ret.State = AllTypesCreateState(e.State)
	ret.Bytes = e.Bytes
	ret.Nilable = OptString{}
	if e.Nilable != nil {
		ret.Nilable.SetTo(*e.Nilable)
	}
	return &ret
}

func NewAllTypesCreates(es []*ent.AllTypes) []AllTypesCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]AllTypesCreate, len(es))
	for i, e := range es {
		r[i] = NewAllTypesCreate(e).Elem()
	}
	return r
}

func (at *AllTypesCreate) Elem() AllTypesCreate {
	if at == nil {
		return AllTypesCreate{}
	}
	return *at
}

func NewAllTypesList(e *ent.AllTypes) *AllTypesList {
	if e == nil {
		return nil
	}
	var ret AllTypesList
	ret.ID = int64(e.ID)
	ret.Int = e.Int
	ret.Int8 = int32(e.Int8)
	ret.Int16 = int32(e.Int16)
	ret.Int32 = e.Int32
	ret.Int64 = e.Int64
	ret.Uint = int64(e.Uint)
	ret.Uint8 = int32(e.Uint8)
	ret.Uint16 = int32(e.Uint16)
	ret.Uint32 = int64(e.Uint32)
	ret.Uint64 = int64(e.Uint64)
	ret.Float32 = e.Float32
	ret.Float64 = e.Float64
	ret.StringType = e.StringType
	ret.Bool = e.Bool
	ret.UUID = e.UUID
	ret.Time = e.Time
	ret.Text = e.Text
	ret.State = AllTypesListState(e.State)
	ret.Bytes = e.Bytes
	ret.Nilable = OptString{}
	if e.Nilable != nil {
		ret.Nilable.SetTo(*e.Nilable)
	}
	return &ret
}

func NewAllTypesLists(es []*ent.AllTypes) []AllTypesList {
	if len(es) == 0 {
		return nil
	}
	r := make([]AllTypesList, len(es))
	for i, e := range es {
		r[i] = NewAllTypesList(e).Elem()
	}
	return r
}

func (at *AllTypesList) Elem() AllTypesList {
	if at == nil {
		return AllTypesList{}
	}
	return *at
}

func NewAllTypesRead(e *ent.AllTypes) *AllTypesRead {
	if e == nil {
		return nil
	}
	var ret AllTypesRead
	ret.ID = int64(e.ID)
	ret.Int = e.Int
	ret.Int8 = int32(e.Int8)
	ret.Int16 = int32(e.Int16)
	ret.Int32 = e.Int32
	ret.Int64 = e.Int64
	ret.Uint = int64(e.Uint)
	ret.Uint8 = int32(e.Uint8)
	ret.Uint16 = int32(e.Uint16)
	ret.Uint32 = int64(e.Uint32)
	ret.Uint64 = int64(e.Uint64)
	ret.Float32 = e.Float32
	ret.Float64 = e.Float64
	ret.StringType = e.StringType
	ret.Bool = e.Bool
	ret.UUID = e.UUID
	ret.Time = e.Time
	ret.Text = e.Text
	ret.State = AllTypesReadState(e.State)
	ret.Bytes = e.Bytes
	ret.Nilable = OptString{}
	if e.Nilable != nil {
		ret.Nilable.SetTo(*e.Nilable)
	}
	return &ret
}

func NewAllTypesReads(es []*ent.AllTypes) []AllTypesRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]AllTypesRead, len(es))
	for i, e := range es {
		r[i] = NewAllTypesRead(e).Elem()
	}
	return r
}

func (at *AllTypesRead) Elem() AllTypesRead {
	if at == nil {
		return AllTypesRead{}
	}
	return *at
}

func NewAllTypesUpdate(e *ent.AllTypes) *AllTypesUpdate {
	if e == nil {
		return nil
	}
	var ret AllTypesUpdate
	ret.ID = int64(e.ID)
	ret.Int = e.Int
	ret.Int8 = int32(e.Int8)
	ret.Int16 = int32(e.Int16)
	ret.Int32 = e.Int32
	ret.Int64 = e.Int64
	ret.Uint = int64(e.Uint)
	ret.Uint8 = int32(e.Uint8)
	ret.Uint16 = int32(e.Uint16)
	ret.Uint32 = int64(e.Uint32)
	ret.Uint64 = int64(e.Uint64)
	ret.Float32 = e.Float32
	ret.Float64 = e.Float64
	ret.StringType = e.StringType
	ret.Bool = e.Bool
	ret.UUID = e.UUID
	ret.Time = e.Time
	ret.Text = e.Text
	ret.State = AllTypesUpdateState(e.State)
	ret.Bytes = e.Bytes
	ret.Nilable = OptString{}
	if e.Nilable != nil {
		ret.Nilable.SetTo(*e.Nilable)
	}
	return &ret
}

func NewAllTypesUpdates(es []*ent.AllTypes) []AllTypesUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]AllTypesUpdate, len(es))
	for i, e := range es {
		r[i] = NewAllTypesUpdate(e).Elem()
	}
	return r
}

func (at *AllTypesUpdate) Elem() AllTypesUpdate {
	if at == nil {
		return AllTypesUpdate{}
	}
	return *at
}

func NewCategoryCreate(e *ent.Category) *CategoryCreate {
	if e == nil {
		return nil
	}
	var ret CategoryCreate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Readonly = NewOptString(e.Readonly)
	return &ret
}

func NewCategoryCreates(es []*ent.Category) []CategoryCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]CategoryCreate, len(es))
	for i, e := range es {
		r[i] = NewCategoryCreate(e).Elem()
	}
	return r
}

func (c *CategoryCreate) Elem() CategoryCreate {
	if c == nil {
		return CategoryCreate{}
	}
	return *c
}

func NewCategoryList(e *ent.Category) *CategoryList {
	if e == nil {
		return nil
	}
	var ret CategoryList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Readonly = NewOptString(e.Readonly)
	return &ret
}

func NewCategoryLists(es []*ent.Category) []CategoryList {
	if len(es) == 0 {
		return nil
	}
	r := make([]CategoryList, len(es))
	for i, e := range es {
		r[i] = NewCategoryList(e).Elem()
	}
	return r
}

func (c *CategoryList) Elem() CategoryList {
	if c == nil {
		return CategoryList{}
	}
	return *c
}

func NewCategoryRead(e *ent.Category) *CategoryRead {
	if e == nil {
		return nil
	}
	var ret CategoryRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Readonly = NewOptString(e.Readonly)
	return &ret
}

func NewCategoryReads(es []*ent.Category) []CategoryRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]CategoryRead, len(es))
	for i, e := range es {
		r[i] = NewCategoryRead(e).Elem()
	}
	return r
}

func (c *CategoryRead) Elem() CategoryRead {
	if c == nil {
		return CategoryRead{}
	}
	return *c
}

func NewCategoryUpdate(e *ent.Category) *CategoryUpdate {
	if e == nil {
		return nil
	}
	var ret CategoryUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Readonly = NewOptString(e.Readonly)
	return &ret
}

func NewCategoryUpdates(es []*ent.Category) []CategoryUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]CategoryUpdate, len(es))
	for i, e := range es {
		r[i] = NewCategoryUpdate(e).Elem()
	}
	return r
}

func (c *CategoryUpdate) Elem() CategoryUpdate {
	if c == nil {
		return CategoryUpdate{}
	}
	return *c
}

func NewCategoryPetsList(e *ent.Pet) *CategoryPetsList {
	if e == nil {
		return nil
	}
	var ret CategoryPetsList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Weight = NewOptInt(e.Weight)
	ret.Birthday = NewOptDateTime(e.Birthday)
	ret.TagID = e.TagID
	ret.Height = OptInt{}
	if e.Height != nil {
		ret.Height.SetTo(*e.Height)
	}
	return &ret
}

func NewCategoryPetsLists(es []*ent.Pet) []CategoryPetsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]CategoryPetsList, len(es))
	for i, e := range es {
		r[i] = NewCategoryPetsList(e).Elem()
	}
	return r
}

func (pe *CategoryPetsList) Elem() CategoryPetsList {
	if pe == nil {
		return CategoryPetsList{}
	}
	return *pe
}

func NewHatCreate(e *ent.Hat) *HatCreate {
	if e == nil {
		return nil
	}
	var ret HatCreate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Type = HatCreateType(e.Type)
	return &ret
}

func NewHatCreates(es []*ent.Hat) []HatCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]HatCreate, len(es))
	for i, e := range es {
		r[i] = NewHatCreate(e).Elem()
	}
	return r
}

func (h *HatCreate) Elem() HatCreate {
	if h == nil {
		return HatCreate{}
	}
	return *h
}

func NewHatList(e *ent.Hat) *HatList {
	if e == nil {
		return nil
	}
	var ret HatList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Type = HatListType(e.Type)
	return &ret
}

func NewHatLists(es []*ent.Hat) []HatList {
	if len(es) == 0 {
		return nil
	}
	r := make([]HatList, len(es))
	for i, e := range es {
		r[i] = NewHatList(e).Elem()
	}
	return r
}

func (h *HatList) Elem() HatList {
	if h == nil {
		return HatList{}
	}
	return *h
}

func NewHatRead(e *ent.Hat) *HatRead {
	if e == nil {
		return nil
	}
	var ret HatRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Type = HatReadType(e.Type)
	return &ret
}

func NewHatReads(es []*ent.Hat) []HatRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]HatRead, len(es))
	for i, e := range es {
		r[i] = NewHatRead(e).Elem()
	}
	return r
}

func (h *HatRead) Elem() HatRead {
	if h == nil {
		return HatRead{}
	}
	return *h
}

func NewHatUpdate(e *ent.Hat) *HatUpdate {
	if e == nil {
		return nil
	}
	var ret HatUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Type = HatUpdateType(e.Type)
	return &ret
}

func NewHatUpdates(es []*ent.Hat) []HatUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]HatUpdate, len(es))
	for i, e := range es {
		r[i] = NewHatUpdate(e).Elem()
	}
	return r
}

func (h *HatUpdate) Elem() HatUpdate {
	if h == nil {
		return HatUpdate{}
	}
	return *h
}

func NewHatWearerRead(e *ent.User) *HatWearerRead {
	if e == nil {
		return nil
	}
	var ret HatWearerRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Age = int64(e.Age)
	ret.Height = NewOptInt64(int64(e.Height))
	ret.FavoriteCatBreed = HatWearerReadFavoriteCatBreed(e.FavoriteCatBreed)
	ret.FavoriteColor = HatWearerReadFavoriteColor(e.FavoriteColor)
	NewOptHatWearerReadFavoriteDogBreed(HatWearerReadFavoriteDogBreed(e.FavoriteDogBreed))
	NewOptHatWearerReadFavoriteFishBreed(HatWearerReadFavoriteFishBreed(e.FavoriteFishBreed))
	return &ret
}

func NewHatWearerReads(es []*ent.User) []HatWearerRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]HatWearerRead, len(es))
	for i, e := range es {
		r[i] = NewHatWearerRead(e).Elem()
	}
	return r
}

func (u *HatWearerRead) Elem() HatWearerRead {
	if u == nil {
		return HatWearerRead{}
	}
	return *u
}

func NewPetCreate(e *ent.Pet) *PetCreate {
	if e == nil {
		return nil
	}
	var ret PetCreate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Weight = NewOptInt(e.Weight)
	ret.Birthday = NewOptDateTime(e.Birthday)
	ret.TagID = e.TagID
	ret.Height = OptInt{}
	if e.Height != nil {
		ret.Height.SetTo(*e.Height)
	}
	ret.Categories = NewPetCreateCategoriesSlice(e.Edges.Categories)
	ret.Owner = NewPetCreateOwner(e.Edges.Owner).Elem()
	return &ret
}

func NewPetCreates(es []*ent.Pet) []PetCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetCreate, len(es))
	for i, e := range es {
		r[i] = NewPetCreate(e).Elem()
	}
	return r
}

func (pe *PetCreate) Elem() PetCreate {
	if pe == nil {
		return PetCreate{}
	}
	return *pe
}

func NewPetCreateCategories(e *ent.Category) *PetCreateCategories {
	if e == nil {
		return nil
	}
	var ret PetCreateCategories
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Readonly = NewOptString(e.Readonly)
	return &ret
}

func NewPetCreateCategoriesSlice(es []*ent.Category) []PetCreateCategories {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetCreateCategories, len(es))
	for i, e := range es {
		r[i] = NewPetCreateCategories(e).Elem()
	}
	return r
}

func (c *PetCreateCategories) Elem() PetCreateCategories {
	if c == nil {
		return PetCreateCategories{}
	}
	return *c
}

func NewPetCreateOwner(e *ent.User) *PetCreateOwner {
	if e == nil {
		return nil
	}
	var ret PetCreateOwner
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Age = int64(e.Age)
	ret.Height = NewOptInt64(int64(e.Height))
	ret.FavoriteCatBreed = PetCreateOwnerFavoriteCatBreed(e.FavoriteCatBreed)
	ret.FavoriteColor = PetCreateOwnerFavoriteColor(e.FavoriteColor)
	NewOptPetCreateOwnerFavoriteDogBreed(PetCreateOwnerFavoriteDogBreed(e.FavoriteDogBreed))
	NewOptPetCreateOwnerFavoriteFishBreed(PetCreateOwnerFavoriteFishBreed(e.FavoriteFishBreed))
	return &ret
}

func NewPetCreateOwners(es []*ent.User) []PetCreateOwner {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetCreateOwner, len(es))
	for i, e := range es {
		r[i] = NewPetCreateOwner(e).Elem()
	}
	return r
}

func (u *PetCreateOwner) Elem() PetCreateOwner {
	if u == nil {
		return PetCreateOwner{}
	}
	return *u
}

func NewPetList(e *ent.Pet) *PetList {
	if e == nil {
		return nil
	}
	var ret PetList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Weight = NewOptInt(e.Weight)
	ret.Birthday = NewOptDateTime(e.Birthday)
	ret.TagID = e.TagID
	ret.Height = OptInt{}
	if e.Height != nil {
		ret.Height.SetTo(*e.Height)
	}
	return &ret
}

func NewPetLists(es []*ent.Pet) []PetList {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetList, len(es))
	for i, e := range es {
		r[i] = NewPetList(e).Elem()
	}
	return r
}

func (pe *PetList) Elem() PetList {
	if pe == nil {
		return PetList{}
	}
	return *pe
}

func NewPetRead(e *ent.Pet) *PetRead {
	if e == nil {
		return nil
	}
	var ret PetRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Weight = NewOptInt(e.Weight)
	ret.Birthday = NewOptDateTime(e.Birthday)
	ret.TagID = e.TagID
	ret.Height = OptInt{}
	if e.Height != nil {
		ret.Height.SetTo(*e.Height)
	}
	return &ret
}

func NewPetReads(es []*ent.Pet) []PetRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetRead, len(es))
	for i, e := range es {
		r[i] = NewPetRead(e).Elem()
	}
	return r
}

func (pe *PetRead) Elem() PetRead {
	if pe == nil {
		return PetRead{}
	}
	return *pe
}

func NewPetUpdate(e *ent.Pet) *PetUpdate {
	if e == nil {
		return nil
	}
	var ret PetUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Weight = NewOptInt(e.Weight)
	ret.Birthday = NewOptDateTime(e.Birthday)
	ret.TagID = e.TagID
	ret.Height = OptInt{}
	if e.Height != nil {
		ret.Height.SetTo(*e.Height)
	}
	return &ret
}

func NewPetUpdates(es []*ent.Pet) []PetUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetUpdate, len(es))
	for i, e := range es {
		r[i] = NewPetUpdate(e).Elem()
	}
	return r
}

func (pe *PetUpdate) Elem() PetUpdate {
	if pe == nil {
		return PetUpdate{}
	}
	return *pe
}

func NewPetCategoriesList(e *ent.Category) *PetCategoriesList {
	if e == nil {
		return nil
	}
	var ret PetCategoriesList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Readonly = NewOptString(e.Readonly)
	return &ret
}

func NewPetCategoriesLists(es []*ent.Category) []PetCategoriesList {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetCategoriesList, len(es))
	for i, e := range es {
		r[i] = NewPetCategoriesList(e).Elem()
	}
	return r
}

func (c *PetCategoriesList) Elem() PetCategoriesList {
	if c == nil {
		return PetCategoriesList{}
	}
	return *c
}

func NewPetFriendsList(e *ent.Pet) *PetFriendsList {
	if e == nil {
		return nil
	}
	var ret PetFriendsList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Weight = NewOptInt(e.Weight)
	ret.Birthday = NewOptDateTime(e.Birthday)
	ret.TagID = e.TagID
	ret.Height = OptInt{}
	if e.Height != nil {
		ret.Height.SetTo(*e.Height)
	}
	return &ret
}

func NewPetFriendsLists(es []*ent.Pet) []PetFriendsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetFriendsList, len(es))
	for i, e := range es {
		r[i] = NewPetFriendsList(e).Elem()
	}
	return r
}

func (pe *PetFriendsList) Elem() PetFriendsList {
	if pe == nil {
		return PetFriendsList{}
	}
	return *pe
}

func NewPetOwnerRead(e *ent.User) *PetOwnerRead {
	if e == nil {
		return nil
	}
	var ret PetOwnerRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Age = int64(e.Age)
	ret.Height = NewOptInt64(int64(e.Height))
	ret.FavoriteCatBreed = PetOwnerReadFavoriteCatBreed(e.FavoriteCatBreed)
	ret.FavoriteColor = PetOwnerReadFavoriteColor(e.FavoriteColor)
	NewOptPetOwnerReadFavoriteDogBreed(PetOwnerReadFavoriteDogBreed(e.FavoriteDogBreed))
	NewOptPetOwnerReadFavoriteFishBreed(PetOwnerReadFavoriteFishBreed(e.FavoriteFishBreed))
	return &ret
}

func NewPetOwnerReads(es []*ent.User) []PetOwnerRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetOwnerRead, len(es))
	for i, e := range es {
		r[i] = NewPetOwnerRead(e).Elem()
	}
	return r
}

func (u *PetOwnerRead) Elem() PetOwnerRead {
	if u == nil {
		return PetOwnerRead{}
	}
	return *u
}

func NewPetRescuerList(e *ent.User) *PetRescuerList {
	if e == nil {
		return nil
	}
	var ret PetRescuerList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Age = int64(e.Age)
	ret.Height = NewOptInt64(int64(e.Height))
	ret.FavoriteCatBreed = PetRescuerListFavoriteCatBreed(e.FavoriteCatBreed)
	ret.FavoriteColor = PetRescuerListFavoriteColor(e.FavoriteColor)
	NewOptPetRescuerListFavoriteDogBreed(PetRescuerListFavoriteDogBreed(e.FavoriteDogBreed))
	NewOptPetRescuerListFavoriteFishBreed(PetRescuerListFavoriteFishBreed(e.FavoriteFishBreed))
	return &ret
}

func NewPetRescuerLists(es []*ent.User) []PetRescuerList {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetRescuerList, len(es))
	for i, e := range es {
		r[i] = NewPetRescuerList(e).Elem()
	}
	return r
}

func (u *PetRescuerList) Elem() PetRescuerList {
	if u == nil {
		return PetRescuerList{}
	}
	return *u
}

func NewUserCreate(e *ent.User) *UserCreate {
	if e == nil {
		return nil
	}
	var ret UserCreate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Age = int64(e.Age)
	ret.Height = NewOptInt64(int64(e.Height))
	ret.FavoriteCatBreed = UserCreateFavoriteCatBreed(e.FavoriteCatBreed)
	ret.FavoriteColor = UserCreateFavoriteColor(e.FavoriteColor)
	NewOptUserCreateFavoriteDogBreed(UserCreateFavoriteDogBreed(e.FavoriteDogBreed))
	NewOptUserCreateFavoriteFishBreed(UserCreateFavoriteFishBreed(e.FavoriteFishBreed))
	return &ret
}

func NewUserCreates(es []*ent.User) []UserCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserCreate, len(es))
	for i, e := range es {
		r[i] = NewUserCreate(e).Elem()
	}
	return r
}

func (u *UserCreate) Elem() UserCreate {
	if u == nil {
		return UserCreate{}
	}
	return *u
}

func NewUserList(e *ent.User) *UserList {
	if e == nil {
		return nil
	}
	var ret UserList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Age = int64(e.Age)
	ret.Height = NewOptInt64(int64(e.Height))
	ret.FavoriteCatBreed = UserListFavoriteCatBreed(e.FavoriteCatBreed)
	ret.FavoriteColor = UserListFavoriteColor(e.FavoriteColor)
	NewOptUserListFavoriteDogBreed(UserListFavoriteDogBreed(e.FavoriteDogBreed))
	NewOptUserListFavoriteFishBreed(UserListFavoriteFishBreed(e.FavoriteFishBreed))
	return &ret
}

func NewUserLists(es []*ent.User) []UserList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserList, len(es))
	for i, e := range es {
		r[i] = NewUserList(e).Elem()
	}
	return r
}

func (u *UserList) Elem() UserList {
	if u == nil {
		return UserList{}
	}
	return *u
}

func NewUserRead(e *ent.User) *UserRead {
	if e == nil {
		return nil
	}
	var ret UserRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Age = int64(e.Age)
	ret.Height = NewOptInt64(int64(e.Height))
	ret.FavoriteCatBreed = UserReadFavoriteCatBreed(e.FavoriteCatBreed)
	ret.FavoriteColor = UserReadFavoriteColor(e.FavoriteColor)
	NewOptUserReadFavoriteDogBreed(UserReadFavoriteDogBreed(e.FavoriteDogBreed))
	NewOptUserReadFavoriteFishBreed(UserReadFavoriteFishBreed(e.FavoriteFishBreed))
	return &ret
}

func NewUserReads(es []*ent.User) []UserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserRead, len(es))
	for i, e := range es {
		r[i] = NewUserRead(e).Elem()
	}
	return r
}

func (u *UserRead) Elem() UserRead {
	if u == nil {
		return UserRead{}
	}
	return *u
}

func NewUserUpdate(e *ent.User) *UserUpdate {
	if e == nil {
		return nil
	}
	var ret UserUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Age = int64(e.Age)
	ret.Height = NewOptInt64(int64(e.Height))
	ret.FavoriteCatBreed = UserUpdateFavoriteCatBreed(e.FavoriteCatBreed)
	ret.FavoriteColor = UserUpdateFavoriteColor(e.FavoriteColor)
	NewOptUserUpdateFavoriteDogBreed(UserUpdateFavoriteDogBreed(e.FavoriteDogBreed))
	NewOptUserUpdateFavoriteFishBreed(UserUpdateFavoriteFishBreed(e.FavoriteFishBreed))
	return &ret
}

func NewUserUpdates(es []*ent.User) []UserUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserUpdate, len(es))
	for i, e := range es {
		r[i] = NewUserUpdate(e).Elem()
	}
	return r
}

func (u *UserUpdate) Elem() UserUpdate {
	if u == nil {
		return UserUpdate{}
	}
	return *u
}

func NewUserAnimalsSavedList(e *ent.Pet) *UserAnimalsSavedList {
	if e == nil {
		return nil
	}
	var ret UserAnimalsSavedList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Weight = NewOptInt(e.Weight)
	ret.Birthday = NewOptDateTime(e.Birthday)
	ret.TagID = e.TagID
	ret.Height = OptInt{}
	if e.Height != nil {
		ret.Height.SetTo(*e.Height)
	}
	return &ret
}

func NewUserAnimalsSavedLists(es []*ent.Pet) []UserAnimalsSavedList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserAnimalsSavedList, len(es))
	for i, e := range es {
		r[i] = NewUserAnimalsSavedList(e).Elem()
	}
	return r
}

func (pe *UserAnimalsSavedList) Elem() UserAnimalsSavedList {
	if pe == nil {
		return UserAnimalsSavedList{}
	}
	return *pe
}

func NewUserBestFriendRead(e *ent.User) *UserBestFriendRead {
	if e == nil {
		return nil
	}
	var ret UserBestFriendRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Age = int64(e.Age)
	ret.Height = NewOptInt64(int64(e.Height))
	ret.FavoriteCatBreed = UserBestFriendReadFavoriteCatBreed(e.FavoriteCatBreed)
	ret.FavoriteColor = UserBestFriendReadFavoriteColor(e.FavoriteColor)
	NewOptUserBestFriendReadFavoriteDogBreed(UserBestFriendReadFavoriteDogBreed(e.FavoriteDogBreed))
	NewOptUserBestFriendReadFavoriteFishBreed(UserBestFriendReadFavoriteFishBreed(e.FavoriteFishBreed))
	return &ret
}

func NewUserBestFriendReads(es []*ent.User) []UserBestFriendRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserBestFriendRead, len(es))
	for i, e := range es {
		r[i] = NewUserBestFriendRead(e).Elem()
	}
	return r
}

func (u *UserBestFriendRead) Elem() UserBestFriendRead {
	if u == nil {
		return UserBestFriendRead{}
	}
	return *u
}

func NewUserFavoriteHatRead(e *ent.Hat) *UserFavoriteHatRead {
	if e == nil {
		return nil
	}
	var ret UserFavoriteHatRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Type = UserFavoriteHatReadType(e.Type)
	return &ret
}

func NewUserFavoriteHatReads(es []*ent.Hat) []UserFavoriteHatRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserFavoriteHatRead, len(es))
	for i, e := range es {
		r[i] = NewUserFavoriteHatRead(e).Elem()
	}
	return r
}

func (h *UserFavoriteHatRead) Elem() UserFavoriteHatRead {
	if h == nil {
		return UserFavoriteHatRead{}
	}
	return *h
}

func NewUserPetsList(e *ent.Pet) *UserPetsList {
	if e == nil {
		return nil
	}
	var ret UserPetsList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Weight = NewOptInt(e.Weight)
	ret.Birthday = NewOptDateTime(e.Birthday)
	ret.TagID = e.TagID
	ret.Height = OptInt{}
	if e.Height != nil {
		ret.Height.SetTo(*e.Height)
	}
	return &ret
}

func NewUserPetsLists(es []*ent.Pet) []UserPetsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserPetsList, len(es))
	for i, e := range es {
		r[i] = NewUserPetsList(e).Elem()
	}
	return r
}

func (pe *UserPetsList) Elem() UserPetsList {
	if pe == nil {
		return UserPetsList{}
	}
	return *pe
}
