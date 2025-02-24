package main

import "fmt"

func (app *App) New_Distributor(name string) error {
	_, distrubutor_exist := app.Distributors[name]
	if distrubutor_exist {
		err := fmt.Errorf("distributor %v already exits", name)
		return err
	}

	app.Distributors[name] = Distributor{
		Name:           name,
		IncludeRegions: make(map[string]bool),
		ExcludeRegions: make(map[string]bool),
		Parent:         nil,
	}
	fmt.Printf("%v addedd successfull\n", name)
	return nil
}

func (app *App) Include_Permissions(name string, place_code string) error {
	distributor, exist := app.Distributors[name]
	if !exist {
		err := fmt.Errorf("no distrubutor with name %v", name)
		return err
	}

	is_place := app.Places.Place_Exist(place_code)
	if !is_place {
		err := fmt.Errorf("invalid place code: %v", place_code)
		return err
	}

	if distributor.Parent != nil {
		if !distributor.Parent.Has_Permission(place_code) {
			err := fmt.Errorf("parent distrubutor has not the permissions")
			return err
		}
		_, exist := distributor.Parent.IncludeRegions[place_code]
		if exist {
			fmt.Println("added permissions to ", name)
			return nil
		}
	}

	_, exist = distributor.ExcludeRegions[place_code]
	if exist {
		delete(distributor.ExcludeRegions, place_code)
	}

	distributor.IncludeRegions[place_code] = true
	fmt.Println("added permissions to ", name)
	return nil
}

func (app *App) Exclude_Permissions(name, place_code string) error {
	d, exist := app.Distributors[name]
	if !exist {
		err := fmt.Errorf("no distrubutor with name %v", name)
		return err
	}

	if !app.Places.Place_Exist(place_code) {
		err := fmt.Errorf("invalid place code: %v", place_code)
		return err
	}

	if d.Parent != nil {
		_, exist = d.Parent.ExcludeRegions[place_code]
		if exist {
			fmt.Println("added exclude permissions to ", name)
			return nil
		}
	}

	_, exist = d.IncludeRegions[place_code]
	if exist {
		delete(d.IncludeRegions, place_code)
	}

	_, exist = d.ExcludeRegions[place_code]
	if !exist {
		d.ExcludeRegions[place_code] = true
	}
	fmt.Println("added exlcude permissions to ", name)
	return nil
}

func (app *App) New_Dist_With_Parent(parent_name, name string) error {
	parent_dist, parent_exist := app.Distributors[parent_name]
	if !parent_exist {
		err := fmt.Errorf("distributor %v does not exits", parent_name)
		return err
	}

	dist, dist_exist := app.Distributors[name]
	if dist_exist {
		if dist.Parent != nil {
			err := fmt.Errorf("%v already has a parent", name)
			return err
		}
		dist.Parent = &parent_dist
		fmt.Printf("%v addedd successfull", name)
		return nil
	}

	app.Distributors[name] = Distributor{
		Name:           name,
		IncludeRegions: make(map[string]bool),
		ExcludeRegions: make(map[string]bool),
		Parent:         &parent_dist,
	}
	fmt.Printf("%v addedd successfull\n", name)
	return nil
}

func (app *App) Get_Permissions(name string) error {
	dist, e := app.Distributors[name]
	if !e {
		err := fmt.Errorf("no distributor with name: %v", name)
		return err
	}

	has_parent := dist.Parent != nil

	fmt.Println("INCLUDE")
	for k := range dist.IncludeRegions {
		app.Places.Get_Name_From_Codes(k)
	}
	if has_parent {
		for k := range dist.Parent.IncludeRegions {
			app.Places.Get_Name_From_Codes(k)
		}
	}
	fmt.Println("==============================")
	fmt.Println("EXCLUDE")
	for k := range dist.ExcludeRegions {
		app.Places.Get_Name_From_Codes(k)
	}
	if has_parent {
		for k := range dist.Parent.ExcludeRegions {
			app.Places.Get_Name_From_Codes(k)
		}
	}
	return nil
}
