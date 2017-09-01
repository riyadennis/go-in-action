package search

//empty struct allocates zero bytes
//used when you need a type but not any state
type defaultMatcher struct {

}
//`m defaultMatcher` is the receiver here
//the method search is now bound to type default matcher
func (m defaultMatcher) Search(feed *Feed, searchTerm string) (r *Result, err error){
	return nil, nil
}

func init(){
	var matcher defaultMatcher
	Register("default", matcher)
}