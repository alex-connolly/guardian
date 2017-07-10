
struct traverser {

};

#define TRAVERSE(name) void traverse_##name(struct traverser* traverser, struct name_node* node)

TRAVERSE(class_decl);
TRAVERSE(interface_decl);
TRAVERSE(contract_decl);
