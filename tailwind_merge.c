#include <php.h>
#include <zend_exceptions.h>

#include "_cgo_export.h"
#include "tailwind_merge.h"
#include "tailwind_merge_arginfo.h"

static int (*original_php_register_internal_extensions_func)(void) = NULL;

ZEND_FUNCTION(tailwind_merge) {
    zval *classes_zval;

    ZEND_PARSE_PARAMETERS_START(1, 1)
        Z_PARAM_ARRAY(classes_zval)
    ZEND_PARSE_PARAMETERS_END();

    HashTable *classes_ht = Z_ARRVAL_P(classes_zval);
    int count = zend_hash_num_elements(classes_ht);

    if (count == 0) {
        RETURN_EMPTY_STRING();
    }

    /* Allocate array of zend_string pointers */
    zend_string **strings = emalloc(sizeof(zend_string *) * count);
    zval *entry;
    int index = 0;

    ZEND_HASH_FOREACH_VAL(classes_ht, entry) {
        if (Z_TYPE_P(entry) != IS_STRING) {
            efree(strings);
            zend_argument_type_error(1, "must be an array of strings, %s given in element %d",
                                     zend_zval_value_name(entry), index);
            RETURN_THROWS();
        }

        strings[index] = Z_STR_P(entry);
        index++;
    }
    ZEND_HASH_FOREACH_END();

    struct go_tailwind_merge_return ret = go_tailwind_merge(strings, count);
    efree(strings);

    if (ret.r0 != NULL) {
        ZVAL_STRING(return_value, ret.r0);
        free(ret.r0);
    } else {
        RETURN_EMPTY_STRING();
    }
}

zend_module_entry ext_module_entry = {
    STANDARD_MODULE_HEADER,
    "tailwind_merge",
    ext_functions,
    NULL, /* MINIT */
    NULL, /* MSHUTDOWN */
    NULL, /* RINIT */
    NULL, /* RSHUTDOWN */
    NULL, /* MINFO */
    "0.1.0",
    STANDARD_MODULE_PROPERTIES
};

PHPAPI int register_internal_extensions(void) {
    if (original_php_register_internal_extensions_func != NULL &&
        original_php_register_internal_extensions_func() != SUCCESS) {
        return FAILURE;
    }

    zend_module_entry *module = &ext_module_entry;
    if (zend_register_internal_module(module) == NULL) {
        return FAILURE;
    }

    return SUCCESS;
}

void register_extension() {
    original_php_register_internal_extensions_func =
        php_register_internal_extensions_func;
    php_register_internal_extensions_func = register_internal_extensions;
}
