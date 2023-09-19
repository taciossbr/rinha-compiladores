from pathlib import Path

def main():
    output_dir = "app/expressions"
    define_ast(
        output_dir,
        "Term",
        [
            ("Binary", "Left Term", "Operation RinOp", "Right Term"),
            ("Bool", "Value bool"),
            ("Integer", "Value int"),
            ("Print",   "Value Term"),
            ("String",  "Value string"),
        ]
    )


def define_ast(output_dir, base_name, terms):
    filename = Path(output_dir) / (base_name.lower() + 's.go')
    with open(filename, 'w') as file:
        print('package expressions', file=file)
        print(file=file)

        # Visitor
        # TODO ver se tem algo melhor que o any, aind anao desisti dos generics
        print("type Visitor interface {", file=file)
        for classname, *_ in terms:
            print(f"\tVisit{classname}(t Rin{classname}) any", file=file)
        print("}", file=file)
        print(file=file)

        for classname, *types in terms:
            # Struct
            print(f"type Rin{classname} struct {{", file=file)
            print(f"\tMeta TermMeta", file=file)
            for _type in types:
                print(f"\t{_type}", file=file)
            print("}", file=file)
            print(file=file)
            # Accept
            print(f"func (t *Rin{classname}) Accept(v Visitor) any {{", file=file)
            print(f"\treturn v.Visit{classname}(*t)", file=file)
            print("}", file=file)
            print(file=file)
            # GetKind
            print(f"func (t *Rin{classname}) GetKind() TermKind {{", file=file)
            print("\treturn t.Meta.Kind", file=file)
            print("}", file=file)
            print(file=file)
            # GetLocation
            print(f"func (t *Rin{classname}) GetLocation() Location {{", file=file)
            print("\treturn t.Meta.Location", file=file)
            print("}", file=file)
            print(file=file)
        
        



if __name__ == '__main__':
    main()
