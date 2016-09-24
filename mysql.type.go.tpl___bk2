{{- $short := (shortname .Name "err" "res" "sqlstr" "db" "XOLog") -}}
{{- $table := (schema .Schema .Table.TableName) -}}
{{- $type := .Name}}
{{- if .Comment -}}
// {{ .Comment }}
{{- else -}}
// {{ .Name }} represents a row from '{{ $table }}'.
{{- end }}

type {{ .Name }} struct {
{{- range .Fields }}
	{{ .Col.ColumnName }} {{ retype .Type }} `json:"{{ .Col.ColumnName }}"` // {{ .Col.ColumnName }} -
{{- end }}
{{- if .PrimaryKey }}

	// xo fields
	_exists, _deleted bool
{{ end }}
}


{{ if .PrimaryKey }}
// Exists determines if the {{ .Name }} exists in the database.
func ({{ $short }} *{{ .Name }}) Exists() bool {
	return {{ $short }}._exists
}

// Deleted provides information if the {{ .Name }} has been deleted from the database.
func ({{ $short }} *{{ .Name }}) Deleted() bool {
	return {{ $short }}._deleted
}

// Insert inserts the {{ .Name }} to the database.
func ({{ $short }} *{{ .Name }}) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if {{ $short }}._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO {{ $table }} (` +
		`{{ colnames .Fields .PrimaryKey.Name }}` +
		`) VALUES (` +
		`{{ colvals .Fields .PrimaryKey.Name }}` +
		`)`

	// run query
	XOLog(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }})
	res, err := db.Exec(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }})
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	{{ $short }}.{{ .PrimaryKey.Name }} = {{ .PrimaryKey.Type }}(id)
	{{ $short }}._exists = true

	return nil
}

// Insert inserts the {{ .Name }} to the database.
func ({{ $short }} *{{ .Name }}) Replace(db XODB) error {
	var err error

	// sql query
	const sqlstr = `REPLACE INTO {{ $table }} (` +
		`{{ colnames .Fields .PrimaryKey.Name }}` +
		`) VALUES (` +
		`{{ colvals .Fields .PrimaryKey.Name }}` +
		`)`

	// run query
	XOLog(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }})
	res, err := db.Exec(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }})
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	{{ $short }}.{{ .PrimaryKey.Name }} = {{ .PrimaryKey.Type }}(id)
	{{ $short }}._exists = true

	return nil
}

// Update updates the {{ .Name }} in the database.
func ({{ $short }} *{{ .Name }}) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !{{ $short }}._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if {{ $short }}._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE {{ $table }} SET ` +
		`{{ colnamesquery .Fields ", " .PrimaryKey.Name }}` +
		` WHERE {{ colname .PrimaryKey.Col }} = ?`

	// run query
	XOLog(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }}, {{ $short }}.{{ .PrimaryKey.Name }})
	_, err = db.Exec(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }}, {{ $short }}.{{ .PrimaryKey.Name }})
	return err
}

// Save saves the {{ .Name }} to the database.
func ({{ $short }} *{{ .Name }}) Save(db XODB) error {
	if {{ $short }}.Exists() {
		return {{ $short }}.Update(db)
	}

	return {{ $short }}.Replace(db)
}

// Delete deletes the {{ .Name }} from the database.
func ({{ $short }} *{{ .Name }}) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !{{ $short }}._exists {
		return nil
	}

	// if deleted, bail
	if {{ $short }}._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM {{ $table }} WHERE {{ colname .PrimaryKey.Col }} = ?`

	// run query
	XOLog(sqlstr, {{ $short }}.{{ .PrimaryKey.Name }})
	_, err = db.Exec(sqlstr, {{ $short }}.{{ .PrimaryKey.Name }})
	if err != nil {
		return err
	}

	// set deleted
	{{ $short }}._deleted = true

	return nil
}

///////////////////////////// Querify gen - ME /////////////////////////////////////////
{{- $deleterType := printf "__%s_Deleter" .Name }}
{{- $updater := printf "__%s_Updater" .Name }}


type {{ $deleterType }} struct {
	wheres   []whereClause
    whereSep string
}

func New{{ $type }}_Deleter()  *{{ $deleterType }} {
    d := {{ $deleterType }} {whereSep: " AND "}
    return &d
}

///////////////// Block ////////////////////////
//varibles init
{{- $colName := "%%%%%%%%%%%" }}
{{ $deleterType := "%%" }}
{{ $ms_gen_types := ms_gen_types() }}

{{block "where" . }}
	
	{{if (eq .Type "int") or (eq .Type "int64") }}
		
		func (u *{{ $deleterType }}){{ $colName }}_In (ins []int) *{{ $deleterType }} {
		    w := whereClause{}
		    var insWhere []interface{}
		    for _, i:= range ins {
		        insWhere = append(insWhere,i)
		    }
		    w.args = insWhere
		    w.condition = " {{ $colName }} IN("+helper.DbQuestionForSqlIn(len(ins))+") "
		    u.wheres = append(u.wheres, w)

		    return u
		}

		func (u *{{$deleterType}}){{ $colName }}_NotIn (ins []int) *{{$deleterType}} {
		    w := whereClause{}
		    var insWhere []interface{}
		    for _, i:= range ins {
		        insWhere = append(insWhere,i)
		    }
		    w.args = insWhere
		    w.condition = " {{ $colName }} NOT IN("+helper.DbQuestionForSqlIn(len(ins))+") "
		    u.wheres = append(u.wheres, w)

		    return u
		}

		{{ $Name := .Name}}
		{{ with $dels }}
			{{ range  .  }}

				func (d *{{$deleterType}}) {{ $colName }}{{ .Suffix }} (val int) *{{$deleterType}} {
				    w := whereClause{}
				    var insWhere []interface{}
				    insWhere = append(insWhere,val)
				    w.args = insWhere
				    w.condition = " {{ $colName }} {{.Condiation}} ? "
				    d.wheres = append(d.wheres, w)
				    	
				    return d
				}

			{{- end }}
		{{- end }}
		
	{{end}}

{{- end }}




////////////// End Block ///////////////////////

/////////////// int - gens - ME ////////////////
{{- range .Fields }}

	{{if eq .Type "string"}}
	func (d *{{ $type }}) {{ .Name }}__STRINGY(db {{.Type}}) {
		d.Params[{{.Name}}] = db
	}
	{{end}}

{{- end }}

{{- $dels := ms_deleter }}
{{- $ms_in := ms_in }}
////////ints
{{- range .Fields }}
	
	{{- $colName = .Col.ColumnName }}

	{{if (eq .Type "int") or (eq .Type "int64") }}
		
		func (u *{{$deleterType}}){{ $colName }}_In (ins []int) *{{$deleterType}} {
		    w := whereClause{}
		    var insWhere []interface{}
		    for _, i:= range ins {
		        insWhere = append(insWhere,i)
		    }
		    w.args = insWhere
		    w.condition = " {{ $colName }} IN("+helper.DbQuestionForSqlIn(len(ins))+") "
		    u.wheres = append(u.wheres, w)

		    return u
		}

		func (u *{{$deleterType}}){{ $colName }}_NotIn (ins []int) *{{$deleterType}} {
		    w := whereClause{}
		    var insWhere []interface{}
		    for _, i:= range ins {
		        insWhere = append(insWhere,i)
		    }
		    w.args = insWhere
		    w.condition = " {{ $colName }} NOT IN("+helper.DbQuestionForSqlIn(len(ins))+") "
		    u.wheres = append(u.wheres, w)

		    return u
		}

		{{ $Name := .Name}}
		{{ with $dels }}
			{{ range  .  }}

				func (d *{{$deleterType}}) {{ $colName }}{{ .Suffix }} (val int) *{{$deleterType}} {
				    w := whereClause{}
				    var insWhere []interface{}
				    insWhere = append(insWhere,val)
				    w.args = insWhere
				    w.condition = " {{ $colName }} {{.Condiation}} ? "
				    d.wheres = append(d.wheres, w)
				    	
				    return d
				}

			{{- end }}
		{{- end }}
		
	{{end}}

{{- end }}

////////strings
{{- range .Fields }}

	{{- $colName := .Col.ColumnName }}

	{{if (eq .Type "string")}}
		
		func (u *{{$deleterType}}){{ $colName }}_In (ins []string) *{{$deleterType}} {
		    w := whereClause{}
		    var insWhere []interface{}
		    for _, i:= range ins {
		        insWhere = append(insWhere,i)
		    }
		    w.args = insWhere
		    w.condition = " {{ $colName }} IN("+helper.DbQuestionForSqlIn(len(ins))+") "
		    u.wheres = append(u.wheres, w)

		    return u
		}

		func (u *{{$deleterType}}){{ $colName }}_NotIn (ins []string) *{{$deleterType}} {
		    w := whereClause{}
		    var insWhere []interface{}
		    for _, i:= range ins {
		        insWhere = append(insWhere,i)
		    }
		    w.args = insWhere
		    w.condition = " {{ $colName }} NOT IN("+helper.DbQuestionForSqlIn(len(ins))+") "
		    u.wheres = append(u.wheres, w)

		    return u
		}

		{{ $Name := .Name}}
		{{ with $dels }}
			{{ range  .  }}

				func (d *{{$deleterType}}) {{ $colName }}{{ .Suffix }} (val string) *{{$deleterType}} {
				    w := whereClause{}
				    var insWhere []interface{}
				    insWhere = append(insWhere,val)
				    w.args = insWhere
				    w.condition = " {{ $colName }} {{.Condiation}} ? "
				    d.wheres = append(d.wheres, w)
				    	
				    return d
				}

			{{- end }}
		{{- end }}
		
	{{end}}

{{- end }}


func (d *{{$deleterType}})Delete (db XODB) (int,error) {
    var err error
    var wheresArr []string
    for _,w := range d.wheres{
        wheresArr = append(wheresArr,d.condition)
    }
    wheresStr := strings.Join(wheresArr, d.whereSep)

    var args []interface{}
    for _,w := range d.wheres{
        args = append(args,w.args...)
    }

    sqlstr := "DELETE FROM {{ $table}} WHERE " + wheresStr

    // run query
    XOLog(sqlstr, args...)
    res, err := db.Exec(sqlstr, args...)
    if err != nil {
        return 0,err
    }

    // retrieve id
    num, err := res.RowsAffected()
    if err != nil {
        return 0,err
    }

    return int(num),nil
}





{{- end }}

